package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"load-management/models"
	"github.com/md96/load-management/db"
	"github.com/md96/load-management/models"
)

func LoginAPI(w http.ResponseWriter, r *http.Request) {
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

func main() {
	db.DBInit()
	/*station := models.Gridstations{
		ID:     1002,
		Status: "Active",
		Name:   "MAH-002",
	}

	result := db.DB.Create(&station)
	if result.Error != nil {
		fmt.Println("Data insertion issue")
	}*/

	fmt.Println("Load management System micro service ")
	http.HandleFunc("/login", LoginAPI)
	fmt.Println("Load Management Service is started on 8082 Port")
	http.ListenAndServe(":8082", nil)
}
