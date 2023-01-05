package rest

import (
	"encoding/json"
	"foo.org/myapp/internal/entities"
	"foo.org/myapp/internal/persistence"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
)

func GetAllSensor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensorType := vars["sensorType"]

	data := persistence.SelectDataSensorFromTo(sensorType)

	if len(data) == 0 {
		http.Error(w, "No data available for this measure type", http.StatusNoContent)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetMoyenneAllDataForADay(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]
	match, _ := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}$", date)
	if !match {
		http.Error(w, "Bad Request, the date must respect the format : YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	allMeasures := persistence.SelectAllDataForADay(date)
	if len(allMeasures) == 0 {
		http.Error(w, "No data available for this day", http.StatusNoContent)
		return
	}

	sensorAve := entities.SensorAvg{}
	for measureNature, measures := range allMeasures {
		average := GetAverage(measures)
		if measureNature == "wind" {
			sensorAve.WindAverage = average
		} else if measureNature == "pressure" {
			sensorAve.PressureAverage = average
		} else {
			sensorAve.TemperatureAverage = average
		}
	}

	jsonData, err := json.Marshal(sensorAve)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetAverage(objects []entities.Sensor) float32 {
	var sum float32 = 0.0
	for _, object := range objects {
		sum += object.Value
	}
	return sum / float32(len(objects))
}
