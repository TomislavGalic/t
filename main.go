package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TomislavGalic/CRUDAPI/models"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func main() {

	godotenv.Load()

	dbURI := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

	DB.AutoMigrate(&models.Vehicle{})

	r := mux.NewRouter()
	r.HandleFunc("/getvehicles", controllers.GetVehicles).Methods("GET")
	r.HandleFunc("/createvehicle", controllers.CreateVehicle).Methods("POST")
	r.HandleFunc("/getvehicle/{id}", controllers.GetVehicle).Methods("GET")
	r.HandleFunc("/updatevehicle/{id}", controllers.UpdateVehicle).Methods("PUT")
	r.HandleFunc("/deletevehicle/{id}", controllers.DeleteVehicle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

/*
func GetVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var vehicle []models.Vehicle
	DB.Find(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func GetVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var vehicle models.Vehicle
	DB.First(&vehicle, params["id"])
	json.NewDecoder(r.Body).Decode(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var vehicle models.Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Create(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var vehicle models.Vehicle
	DB.First(&vehicle, params["id"])
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Save(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var vehicle models.Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Delete(&vehicle, params["id"])
	json.NewEncoder(w).Encode("The user is deleted")
}
*/
