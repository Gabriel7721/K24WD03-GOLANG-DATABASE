package main

import (
	configs "SliceServer/Configs"
	routers "SliceServer/Routers"
	"fmt"
	"net/http"
)

func main() {
	configs.MongoDBConnect()
	configs.MySQLConnect()

	// services.LoadUsers()

	routers.UserRouter()
	routers.OrderRouter()

	fmt.Println("Server is running at http://localhost:9999")
	http.ListenAndServe(":9999", nil)
}
