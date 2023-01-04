package rest

import (
	"encoding/json"
	"fmt"
	"foo.org/myapp/internal/persistence"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAllSensor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensorType := vars["sensorType"]

	data := persistence.SelectDataSensorFromTo(sensorType)
	fmt.Println(sensorType)
	fmt.Println(data)

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
