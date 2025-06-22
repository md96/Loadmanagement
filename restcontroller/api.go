package restcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"load-management/models"
	"github.com/md96/load-management/db"
	"github.com/md96/load-management/models"
)

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

func Stationcreation(w http.ResponseWriter, r *http.Request) {
	var station models.Gridstations

	if r.Method == http.MethodPost {

		json.NewDecoder(r.Body).Decode(&station)
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
