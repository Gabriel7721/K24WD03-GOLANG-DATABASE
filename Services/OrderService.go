package services

import (
	configs "SliceServer/Configs"
	models "SliceServer/Models"
)

func CreateOrder(order models.Order) int64 {
	result, _ := configs.MySQLDB.Exec("INSERT INTO orders (user_id,product_name,quantity,price) VALUES (?,?,?,?)",
		order.UserID, order.ProductName, order.Quantity, order.Price)
	id, _ := result.LastInsertId()
	return id
}
func GetAllOrders()     {}
func GetOrderByID()     {}
func GetOrderByUserID() {}
func UpdateOrderByID()  {}
func DeletedOrderByID() {}
