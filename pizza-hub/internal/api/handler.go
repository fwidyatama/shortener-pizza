package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pizza-hub/internal/model"
	"pizza-hub/internal/response"
	"pizza-hub/internal/transport"
	"sync"
	"time"
)

var (
	chefs     []model.Chef
	orderChan = make(chan model.Order, 20)
	mutex     sync.Mutex

	orderID = 1

	pizzaMenu = map[string]model.Pizza{
		"pizza-cheese": {
			Name:         "Pizza Cheese",
			CookDuration: 3 * time.Second,
		},
		"pizza-bbq": {
			Name:         "Pizza BBQ",
			CookDuration: 5 * time.Second,
		},
	}
)

type apiHandler struct{}

type ApiHandler interface {
	GetMenu(w http.ResponseWriter, r *http.Request)
	AddChef(w http.ResponseWriter, r *http.Request)
	AddOrder(w http.ResponseWriter, r *http.Request)
}

func NewHandler() ApiHandler {
	return &apiHandler{}
}

func (a *apiHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	pizza := make([]*transport.PizzaRes, 0)

	for _, p := range pizzaMenu {
		newPizza := &transport.PizzaRes{
			Name:         p.Name,
			CookDuration: fmt.Sprintf("%.0f second", p.CookDuration.Seconds()),
		}
		pizza = append(pizza, newPizza)
	}
	response.SuccessResponse(w, pizza)
}

func (a *apiHandler) AddChef(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	req := transport.ChefReq{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errMsg := "Error when processing request"
		response.ErrorResponse(w, errMsg, http.StatusBadRequest)
		return
	}

	newChef := model.Chef{
		ID:        len(chefs) + 1,
		Name:      req.Name,
		IsWorking: false,
	}
	chefs = append(chefs, newChef)

	go assignToCook(newChef.ID)

	response.SuccessResponse(w, transport.MapToChefResponse(&newChef))

	// distribute order to new chef if for when in the middle of queue new chef is added
	distributeOrderToNewChef(newChef)
}

func (a *apiHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	orders := make([]*model.Order, 0)
	req := make([]transport.OrderReq, 0)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errMsg := "Error when processing request"
		response.ErrorResponse(w, errMsg, http.StatusBadRequest)
		return
	}

	for _, orderReq := range req {
		wg.Add(1)

		go func(orderReq transport.OrderReq) {
			defer wg.Done()

			order := model.Order{
				ID:        orderID,
				Customer:  orderReq.Name,
				PizzaType: orderReq.Type,
			}

			orderChan <- order

			orderID++

			orders = append(orders, &order)

			currentTime := time.Now().Format(time.DateTime)
			fmt.Printf("Order for customer => %s (%s) already received at %s \n", order.Customer, order.PizzaType, currentTime)

		}(orderReq)
	}

	wg.Wait()
	response.SuccessResponse(w, transport.MapToOrderResponse(orders))
}
