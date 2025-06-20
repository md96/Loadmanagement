package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("Login Api Requested ")
		w.Header().Set("Content-Type", "application/json")
		//fmt.Fprintf(w, "Login API Requested")
		response := map[string]string{
			"Message": "Login API requested",
		}
		json.NewEncoder(w).Encode(response)
	}
}

func main() {

	fmt.Println("Load management System micro service ")
	http.HandleFunc("/login", LoginAPI)
	fmt.Println("Load Management Service is started on 8082 Port")
	http.ListenAndServe(":8082", nil)
}
