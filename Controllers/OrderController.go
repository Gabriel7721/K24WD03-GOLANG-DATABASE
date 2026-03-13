package controllers

import (
	models "SliceServer/Models"
	services "SliceServer/Services"
	"encoding/json"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)

	if !services.UserExistsById(order.UserID) {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{
		"Message":  "Order created successfully.",
		"Order ID": services.CreateOrder(order),
	})
}
