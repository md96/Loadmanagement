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
