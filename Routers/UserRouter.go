package routers

import (
	controllers "SliceServer/Controllers"
	"fmt"
	"net/http"
)

func UserRouter() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			controllers.GetAllUsers(w, r)
			return
		}
		controllers.CreateNewUser(w, r)
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetUserById(w, r)
		// case "PUT":
		// 	controllers.UpdateUserById(w, r)
		// case "DELETE":
		// 	controllers.DeleteUserById(w, r)
		default:
			fmt.Println("Invalid Method")
		}
	})
}
