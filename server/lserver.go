package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Client struct {
	id   int    "json:'id:9984'"
	name string "json:'name:janez'"
	age  int    "json:age:27"
}

var people []Client

func getPersonsEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.router.NewRouter()
	router.HandleFun("/clients", getPersonsEndpoint().Methods("GET"))

}
