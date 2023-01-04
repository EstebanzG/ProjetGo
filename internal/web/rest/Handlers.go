package rest

import (
	"encoding/json"
	"fmt"
	"foo.org/myapp/internal/entities"
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

func GetMoyenneAllDataForADay(w http.ResponseWriter, r *http.Request) {
	winds, temperatures, pressures := persistence.SelectAllDataForADay()
	windMoyenne := GetAverage(winds)
	temperatureMoyenne := GetAverage(temperatures)
	pressureMoyenne := GetAverage(pressures)
	fmt.Println(windMoyenne)
	fmt.Println(temperatureMoyenne)
	fmt.Println(pressureMoyenne)

	sensorMoy := entities.SensorAvg{
		WindAverage:        windMoyenne,
		TemperatureAverage: temperatureMoyenne,
		PressureAverage:    pressureMoyenne,
	}
	jsonData, err := json.Marshal(sensorMoy)
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
