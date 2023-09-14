package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"sync"
//	"time"
//)
//
//// Chef represents a chef who can prepare orders.
//type Chef struct {
//	ID   int
//	Name string
//}
//
//// MenuItem represents a pizza menu item.
//type MenuItem struct {
//	Name           string
//	DurationToCook time.Duration
//}
//
//// Order represents a customer's pizza order.
//type Order struct {
//	ID       int
//	Customer string
//	Pizza    MenuItem
//	Chef     *Chef
//	Complete bool
//}
//
//var (
//	chefs     []*Chef
//	menuItems = []MenuItem{
//		{"Pizza Cheese", 3 * time.Second},
//		{"Pizza BBQ", 5 * time.Second},
//	}
//	orders  []*Order
//	mutex   sync.Mutex
//	orderID = 1
//	chefID  = 1
//)
//
//func main() {
//	r := http.NewServeMux()
//	r.HandleFunc("/chef", AddChef)
//	r.HandleFunc("/menus", ListMenus)
//	r.HandleFunc("/orders", PlaceOrder)
//
//	fmt.Println("PizzaHub API is running on :8080...")
//	http.ListenAndServe(":8080", r)
//}
//
//// AddChef adds a new chef to the chef pool.
//func AddChef(w http.ResponseWriter, r *http.Request) {
//	mutex.Lock()
//	defer mutex.Unlock()
//
//	chefName := r.FormValue("name")
//	if chefName == "" {
//		http.Error(w, "Chef name is required", http.StatusBadRequest)
//		return
//	}
//
//	chef := &Chef{
//		ID:   chefID,
//		Name: chefName,
//	}
//
//	// Assign the new chef to serve any available orders.
//	for _, order := range orders {
//		if order.Chef == nil {
//			order.Chef = chef
//			break
//		}
//	}
//
//	chefs = append(chefs, chef)
//	chefID++
//
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(chef)
//}
//
//// ListMenus returns the list of available pizza menu items.
//func ListMenus(w http.ResponseWriter, r *http.Request) {
//	json.NewEncoder(w).Encode(menuItems)
//}
//
//// PlaceOrder allows customers to place a pizza order.
//func PlaceOrder(w http.ResponseWriter, r *http.Request) {
//	mutex.Lock()
//	defer mutex.Unlock()
//
//	pizzaName := r.FormValue("pizza")
//	customerName := r.FormValue("customer")
//	if pizzaName == "" || customerName == "" {
//		http.Error(w, "Pizza name and customer name are required", http.StatusBadRequest)
//		return
//	}
//
//	var selectedPizza MenuItem
//	for _, item := range menuItems {
//		if item.Name == pizzaName {
//			selectedPizza = item
//			break
//		}
//	}
//
//	if selectedPizza.Name == "" {
//		http.Error(w, "Pizza not found in the menu", http.StatusNotFound)
//		return
//	}
//
//	// Find an available chef to prepare the order.
//	var chef *Chef
//	for _, c := range chefs {
//		if c.IsAvailable() {
//			chef = c
//			break
//		}
//	}
//
//	if chef == nil {
//		http.Error(w, "No available chefs to prepare the order", http.StatusServiceUnavailable)
//		return
//	}
//
//	order := &Order{
//		ID:       orderID,
//		Customer: customerName,
//		Pizza:    selectedPizza,
//		Chef:     chef,
//	}
//
//	orders = append(orders, order)
//	orderID++
//
//	go order.Prepare()
//
//	//RemoveCompletedOrders()
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(order)
//}
//
//// IsAvailable checks if a chef is available to take an order.
//func (c *Chef) IsAvailable() bool {
//	for _, order := range orders {
//		if order.Chef != nil && order.Chef.ID == c.ID && !order.Complete {
//			return false
//		}
//	}
//	return true
//}
//
//// Prepare simulates the process of preparing a pizza by a chef.
//func (o *Order) Prepare() {
//	time.Sleep(o.Pizza.DurationToCook)
//	mutex.Lock()
//	o.ProcessOrder(*o)
//	defer mutex.Unlock()
//	o.Complete = true
//}
//
//func (o *Order) ProcessOrder(order Order) {
//	startTime := time.Now().Format(time.DateTime)
//	fmt.Printf("Your order is processed by chef %s at %s\n", order.Chef.Name, startTime)
//	time.Sleep(order.Pizza.DurationToCook)
//	endTime := time.Now().Format(time.DateTime)
//	fmt.Printf("Completed order for %s: %s finish in %s \n", order.Customer, order.Pizza.Name, endTime)
//
//}
