package rest

import (
	"encoding/json"
	"fmt"
	"foo.org/myapp/internal/entities"
	"foo.org/myapp/internal/persistence"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func GetBySensorType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensorType := vars["sensorType"]

	data := persistence.SelectByType(sensorType)

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func format_date(date string) (time.Time, error) {
	index := strings.LastIndex(date, "-")
	byte := []byte(date)
	byte[index] = ' '
	date = string(byte)

	return time.Parse("2000-01-01 10:10:10", date)
}

func between_date(start, end time.Time) []time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, start.Hour(), 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, start.Hour(), 0, 0, 0, time.UTC)

	var res []time.Time

	for start.Before(end) {
		res = append(res, start)
		start = start.Add(time.Hour)
	}
	res = append(res, end)
	return res
}

func GetBySensorTypeBetweenDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensorType := vars["sensorType"]
	date1 := vars["date1"]
	date2 := vars["date2"]
	match2, _ := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2}:[0-9]{2}:[0-9]{2}$", date1)
	match1, _ := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2}:[0-9]{2}:[0-9]{2}$", date2)
	if !match1 || !match2 {
		http.Error(w, "Bad Request, the date must respect the format : YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	date1_time, err1 := format_date(date1)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	date2_time, err2 := format_date(date2)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	if date1_time.After(date2_time) {
		tmp := date1_time
		date1_time = date2_time
		date2_time = tmp
	}

	allDates := between_date(date1_time, date2_time)
	data := persistence.SelectAllSensorTypeDateHour(sensorType, allDates)
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
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetMoyenneAllDataForADay(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airport, exist := vars["airportId"]
	if exist {
		match, _ := regexp.MatchString("^[A-Z]{3}$", airport)
		if !match {
			http.Error(w, "Bad Request, the airport IATA code is invalid", 400)
			return
		}
	} else {
		airport = "*"
	}

	date := vars["date"]
	match, _ := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}$", date)
	if !match {
		http.Error(w, "Bad Request, the date must respect the format : YYYY-MM-DD", 401)
		return
	}
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		http.Error(w, "Bad Request, the day isn't exist", 402)
		return
	}

	allMeasures := persistence.SelectAllDataForADay(airport, date)
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
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetAverage(objects []entities.SensorValue) float32 {
	var sum float32 = 0.0
	for _, object := range objects {
		sum += object.Value
	}
	return sum / float32(len(objects))
}
