package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var user User
	params := mux.Vars(r)

	user = User{
		ID:   params["id"],
		Name: "FullName",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{id:[0-9]+}", getUser).Methods("GET")
	http.ListenAndServe(":80", router)
}
