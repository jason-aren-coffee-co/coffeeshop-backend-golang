package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jason-gill00/coffee-shop-backend-golang/src/api/database"
)

// type Orders struct {
// 	CoffeeOrders []database.Order `json:"coffee_orders"`
// }

func GetHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.Connect()
	orders := database.GetOrderHistory(db)
	json.NewEncoder(w).Encode(orders)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	var orders []database.Order
	err := json.NewDecoder(r.Body).Decode(&orders)
	if err != nil {
		panic(err)
	}
	fmt.Println(orders)
	orderBytes, _ := json.MarshalIndent(orders, "", "  ")
	fmt.Println(string(orderBytes))
	database.CreateOrder(orders, db)

}
