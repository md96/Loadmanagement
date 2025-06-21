package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"load-management/models"
	"github.com/md96/load-management/db"
	"github.com/md96/load-management/models"
)

func stationcreation(w http.ResponseWriter, r *http.Request) {
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

func getallstations(w http.ResponseWriter, r *http.Request) {

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

func main() {

	//DB Connection Initialization
	db.DBInit()

	fmt.Println("Load management System micro service ")
	http.HandleFunc("/stationcreation", stationcreation)
	http.HandleFunc("/getallstations", getallstations)
	fmt.Println("Load Management Service is started on 8082 Port")
	http.ListenAndServe(":8082", nil)
}
