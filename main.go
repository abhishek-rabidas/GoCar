package main

import (
	"gocar/config"
	"gocar/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.InitializeDB()
	router := mux.NewRouter()

	router.HandleFunc("/car", services.AddCar).Methods("POST")
	router.HandleFunc("/car/{id}", services.GetCarById).Methods("GET")
	router.HandleFunc("/car", services.GetAllCars).Methods("GET")
	router.HandleFunc("/car", services.UpdateCar).Methods("PUT")
	router.HandleFunc("/car", services.DeleteCar).Methods("DELETE")
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":4000", router))

}
