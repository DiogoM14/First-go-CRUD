package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	Id    int
	Make  string
	Model string
	Price int
}

var vehicles = []Vehicle{
	{1, "Ford", "Fiesta", 23000},
	{2, "Renault", "Clio", 5000},
	{3, "Honda", "Civic", 50000},
	{4, "Ford", "Mustang", 35000},
}

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/cars", returnAllCars).Methods("GET")
	http.ListenAndServe(":8080", router)
}
