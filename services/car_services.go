package services

import (
	"encoding/json"
	"fmt"
	"gocar/config"
	"gocar/models"
	"net/http"

	"github.com/gorilla/mux"
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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var car models.Car
	DB.First(&car, params["id"])
	json.NewEncoder(w).Encode(car)
}

func GetAllCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cars []models.Car
	DB.Find(&cars)
	json.NewEncoder(w).Encode(cars)

}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var car models.Car
	DB.First(&car, params["id"])
	json.NewDecoder(r.Body).Decode(&car)
	DB.Save(&car)
	json.NewEncoder(w).Encode(car)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var car models.Car
	DB.Delete(&car, params["id"])
	json.NewEncoder(w).Encode("The Car is Deleted Successfully!")
}
