package services

import (
	"encoding/json"
	"fmt"
	"gocar/config"
	"gocar/models"
	"net/http"

	"gorm.io/gorm"
)

var DB *gorm.DB = config.InitializeDB()

func AddCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var car models.Car

	json.NewDecoder(r.Body).Decode(&car)
	fmt.Printf("%#v\n", car)
	DB.Create(&car)

	json.NewEncoder(w).Encode(car)

}

func GetCarById(w http.ResponseWriter, r *http.Request) {

}

func GetAllCars(w http.ResponseWriter, r *http.Request) {

}

func UpdateCar(w http.ResponseWriter, r *http.Request) {

}

func DeleteCar(w http.ResponseWriter, r *http.Request) {

}
