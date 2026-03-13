package controllers

import (
	models "SliceServer/Models"
	services "SliceServer/Services"
	"encoding/json"
	"net/http"
	"strings"
)

func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(services.GetAllUsers())
}
func CreateNewUser(res http.ResponseWriter, req *http.Request) {
	var new models.User
	json.NewDecoder(req.Body).Decode(&new)
	json.NewEncoder(res).Encode(services.Create(new))
}

func GetUserById(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/users/")
	u := services.GetUserById(id)
	json.NewEncoder(res).Encode(u)
}

func UpdateUserById(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/users/")
	var update models.User
	json.NewDecoder(req.Body).Decode(&update)
	u := services.UpdateUserById(id, update)
	json.NewEncoder(res).Encode(u)
}

func DeleteUserById(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/users/")
	services.DeleteUser(id)
}
