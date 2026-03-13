package main

import (
	configs "SliceServer/Configs"
	routers "SliceServer/Routers"
	"fmt"
	"net/http"
)

func main() {
	configs.MongoDBConnect()
	// services.LoadUsers()
	routers.UserRouter()
	fmt.Println("Server is running at http://localhost:9999")
	http.ListenAndServe(":9999", nil)
}
