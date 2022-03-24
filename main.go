package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	Id    int
	Make  string
	Model string
	Price int
}

var vehicles = []Vehicle{
	{1, "ford", "Fiesta", 23000},
	{2, "renault", "Clio", 5000},
	{3, "honda", "Civic", 50000},
	{4, "ford", "Mustang", 35000},
}

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func returnCarsByBrand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carM := strings.ToLower(vars["make"])

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
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range vehicles {
		if v.Id == carId {
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)

}

func createCar(w http.ResponseWriter, r *http.Request) {
	var newCar Vehicle
	json.NewDecoder(r.Body).Decode(&newCar)
	vehicles = append(vehicles, newCar)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func removeCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])

	// cars := &[]Vehicle{}

	if err != nil {
		fmt.Println(err)
	}

	for _, car := range vehicles {
		if car.Id == carId {
			vehicles = append(vehicles[:car.Id], vehicles[car.Id+1:]...)
		}
	}

	json.NewEncoder(w).Encode(vehicles)
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
