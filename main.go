package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"

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

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0 << 50)

	file, handler, err := r.FormFile("File")
	if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
	defer file.Close()

	tempFile, err := ioutil.TempFile("", handler.Filename)
    if err != nil {
        fmt.Println(err)
    }
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }

	ioutil.WriteFile(handler.Filename, fileBytes, 0644)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{id:[0-9]+}", getUser).Methods("GET")
	router.HandleFunc("/upload", uploadFile).Methods("POST")
	router.PathPrefix("/download/").Handler(http.StripPrefix("/download", http.FileServer(http.Dir(""))))
	http.ListenAndServe(":80", router)
}
