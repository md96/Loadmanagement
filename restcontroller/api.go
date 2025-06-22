package restcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"load-management/models"
	"github.com/md96/load-management/db"
	"github.com/md96/load-management/models"
)

// Deletestation godoc
// @Summary Delete a station
// @Description Delete station by ID
// @Tags Stations
// @Accept  json
// @Produce  json
// @Param id query string true "Station ID"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid ID"
// @Router /deletestation [delete]
func Deletestation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete station  API is called ")
	query := r.URL.Query()
	id := query.Get("id")
	if id == "" {
		http.Error(w, "Station Id is empty", http.StatusBadRequest)
		return
	}

	result := db.DB.Delete(&models.Gridstations{}, "id=?", id)
	if result.Error != nil {
		http.Error(w, "There is issue while deletng data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"Message": fmt.Sprintf("Station %s deleted succefully", id)})

}

// Stationcreation godoc
// @Summary Create a new station
// @Description Create a new grid station entry
// @Tags Stations
// @Accept  json
// @Produce  json
// @Param station body models.Gridstations true "Station Data"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stationcreation [post]
func Stationcreation(w http.ResponseWriter, r *http.Request) {
	var station models.Gridstations

	if r.Method == http.MethodPost {

		json.NewDecoder(r.Body).Decode(&station)

		fmt.Println("Requested data", station)

		result := db.DB.Create(&station)
		if result.Error != nil {
			fmt.Println("Data insertion issue")
		}

		fmt.Println("Station creation API is called ")
		w.Header().Set("Content-Type", "application/json")
		//fmt.Fprintf(w, "Login API Requested")
		response := map[string]string{
			"Message": "Grid stations data inserted succefully",
		}
		json.NewEncoder(w).Encode(response)
	}

}

// Getallstations godoc
// @Summary Get all stations
// @Description Retrieve all grid stations
// @Tags Stations
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Gridstations
// @Router /getallstations [get]
func Getallstations(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all stations API is called ")

	var stations []models.Gridstations
	if r.Method == http.MethodGet {
		result := db.DB.Find(&stations)
		if result.Error != nil {
			http.Error(w, "Error while fetching data", http.StatusInternalServerError)
			return
		}

	}
	json.NewEncoder(w).Encode(stations)

}
