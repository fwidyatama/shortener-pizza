package transport

import "pizza-hub/internal/model"

type PizzaRes struct {
	Name         string `json:"name"`
	CookDuration string `json:"cook_duration"`
}

type OrderRes struct {
	ID          int    `json:"id"`
	Customer    string `json:"customer"`
	PizzaType   string `json:"pizza_type"`
	IsProcessed bool   `json:"is_processed"`
}
type ChefRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func MapToOrderResponse(orders []*model.Order) []*OrderRes {
	orderResp := make([]*OrderRes, 0)

	for _, order := range orders {
		temp := &OrderRes{
			ID:          order.ID,
			Customer:    order.Customer,
			PizzaType:   order.PizzaType,
			IsProcessed: true,
		}
		orderResp = append(orderResp, temp)
	}

	return orderResp
}

func MapToChefResponse(c *model.Chef) *ChefRes {
	chef := &ChefRes{
		ID:   c.ID,
		Name: c.Name,
	}

	return chef
}
