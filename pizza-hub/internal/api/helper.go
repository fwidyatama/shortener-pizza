package api

import (
	"fmt"
	"pizza-hub/internal/model"
	"strings"
	"time"
)

func assignToCook(chefID int) {
	for {
		order, ok := <-orderChan
		if !ok {
			return
		}

		currentChef := chefs[chefID-1]
		currentChef.IsWorking = true

		pizza := getPizzaByName(order.PizzaType)
		duration := pizza.CookDuration

		time.Sleep(duration)
		notifyFinishCook(order, currentChef)

		currentChef.IsWorking = false
	}
}

func distributeOrderToNewChef(newChef model.Chef) {
	for {
		select {
		case order, ok := <-orderChan:
			if !ok {
				return
			}

			if !newChef.IsWorking {
				go assignToCook(newChef.ID)
				newChef.IsWorking = true

				notifyFinishCook(order, newChef)

				newChef.IsWorking = false
			} else {
				orderChan <- order
				break
			}
		default:
			return
		}
	}
}

func notifyFinishCook(order model.Order, chef model.Chef) {
	currentTime := time.Now().Format(time.DateTime)
	fmt.Printf("Order for customer => %s (%s) already cooked by %s finish at %s \n", order.Customer, order.PizzaType, chef.Name, currentTime)
}

func getPizzaByName(name string) *model.Pizza {
	name = strings.ToLower(name)
	name = strings.Replace(name, " ", "-", 1)
	pizza, isFound := pizzaMenu[name]
	if !isFound {
		return nil
	}

	return &pizza
}
