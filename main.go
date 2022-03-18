package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func returnCarsByBrand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carM := vars["make"]

	cars := &[]Vehicle{}

	for _, car := range vehicles {
		if car.Make == carM {
			*cars = append(*cars, car)
		}
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cars)
}

func returnCarsById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Println(err)
	}

	for _, car := range vehicles {
		if car.Id == carId {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(car)
		}
	}
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

func createCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

func removeCarById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/cars", returnAllCars).Methods("GET")
	router.HandleFunc("/cars/make/{make}", returnCarsByBrand).Methods("GET")
	router.HandleFunc("/cars/{id}", returnCarsById).Methods("GET")
	router.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	router.HandleFunc("/cars", createCar).Methods("POST")
	router.HandleFunc("/cars/{id}", removeCarById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
