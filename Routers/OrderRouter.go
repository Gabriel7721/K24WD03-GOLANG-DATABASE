package routers

import (
	controllers "SliceServer/Controllers"
	"net/http"
)

func OrderRouter() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			controllers.CreateOrder(w, r)
			return
		}
		// controllers.CreateNewUser(w, r)
	})

	// http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case "GET":
	// 		controllers.GetUserById(w, r)
	// 	case "PUT":
	// 		controllers.UpdateUserById(w, r)
	// 	case "DELETE":
	// 		controllers.DeleteUserById(w, r)
	// 	default:
	// 		// fmt.Println("Invalid Method")
	// 		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	// 	}
	// })
}
