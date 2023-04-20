package main

import (
	"gocar/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/car", services.AddCar).Methods("POST")
	router.HandleFunc("/car/{id}", services.GetCarById).Methods("GET")
	router.HandleFunc("/car", services.GetAllCars).Methods("GET")
	router.HandleFunc("/car", services.UpdateCar).Methods("PUT")
	router.HandleFunc("/car", services.DeleteCar).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1234", router))

}
