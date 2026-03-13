package services

import (
	configs "SliceServer/Configs"
	models "SliceServer/Models"
)

func CreateOrder(order models.Order) {
	configs.MySQLDB.Exec("INSERT INTO orders (user_id,product_name,quantity,price) VALUES (?,?,?,?)")
}
func GetAllOrders()     {}
func GetOrderByID()     {}
func GetOrderByUserID() {}
func UpdateOrderByID()  {}
func DeletedOrderByID() {}
