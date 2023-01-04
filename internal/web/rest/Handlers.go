package rest

import (
	"encoding/json"
	"foo.org/myapp/internal/persistence"
	"net/http"
)

func GetAllSensor(w http.ResponseWriter, r *http.Request) {
	data := persistence.SelectDataSensorFromTo("wind")
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
