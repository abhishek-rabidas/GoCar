package services

import (
	"gocar/config"
	"net/http"
)

func AddCar(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	requestBody := r.Body

}

func GetCarById(w http.ResponseWriter, r *http.Request) {

}

func GetAllCars(w http.ResponseWriter, r *http.Request) {

}

func UpdateCar(w http.ResponseWriter, r *http.Request) {

}

func DeleteCar(w http.ResponseWriter, r *http.Request) {

}
