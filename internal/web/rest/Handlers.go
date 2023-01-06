package rest

import (
	"encoding/json"
	"fmt"
	"foo.org/myapp/internal/entities"
	"foo.org/myapp/internal/persistence"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"
)

// ---------------------- route function ----------------------

func GetByTypeBetweenDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensorType := vars["sensorType"]
	date1 := vars["date1"]
	date2 := vars["date2"]
	airportIATA, exist := vars["airportIATA"]
	if exist {
		match, _ := regexp.MatchString("^[A-Z]{3}$", airportIATA)
		if !match {
			http.Error(w, "Bad Request, the airport IATA code is invalid", 400)
			return
		}
	} else {
		airportIATA = "*"
	}

	match1, _ := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2}:[0-9]{2}$", date1)
	match2, _ := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2}:[0-9]{2}$", date2)
	if !match1 || !match2 {
		http.Error(w, "Bad Request, the date must respect the format : YYYY-MM-DD-hh:mm", http.StatusBadRequest)
		return
	}

	date1_time, err1 := formatDate(date1)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	date2_time, err2 := formatDate(date2)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

	if date1_time.After(date2_time) {
		http.Error(w, "Bad Request, change date order", 400)
		return
	}

	var keysOfDate []string
	for date := date1_time; date.Before(date2_time) || date.Equal(date2_time); date = date.AddDate(0, 0, 1) {
		keysOfDate = append(keysOfDate, persistence.SelectKeysByDate(airportIATA, sensorType, date.Format("2006-01-02"))...)
	}

	allKeys := keysBetweenDate(date1_time, date2_time, keysOfDate)
	sort.Strings(allKeys)

	if len(allKeys) == 0 {
		http.Error(w, "No data available for this measure type", http.StatusNoContent)
		return
	}
	data := persistence.GetForKeys(allKeys)
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
		log.Fatal(err)
		return
	}
}

func GetByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensorType := vars["sensorType"]
	airportIATA, exist := vars["airportIATA"]
	if exist {
		match, _ := regexp.MatchString("^[A-Z]{3}$", airportIATA)
		if !match {
			http.Error(w, "Bad Request, the airport IATA code is invalid", 400)
			return
		}
	} else {
		airportIATA = "*"
	}

	allKeys := persistence.SelectKeys(airportIATA, sensorType)
	sort.Strings(allKeys)
	data := persistence.GetForKeys(allKeys)

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetAverageByDay(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	airportIATA, exist := vars["airportIATA"]
	if exist {
		match, _ := regexp.MatchString("^[A-Z]{3}$", airportIATA)
		if !match {
			http.Error(w, "Bad Request, the airport IATA code is invalid", 400)
			return
		}
	} else {
		airportIATA = "*"
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

	allMeasures := persistence.SelectAllDataForADay(airportIATA, date)
	if len(allMeasures) == 0 {
		http.Error(w, "No data available for this day", http.StatusNoContent)
		return
	}

	measuresAve := entities.MeasureAvg{}
	for measureNature, measures := range allMeasures {
		average := average(measures)
		if measureNature == "wind" {
			measuresAve.WindAverage = average
		} else if measureNature == "pressure" {
			measuresAve.PressureAverage = average
		} else {
			measuresAve.TemperatureAverage = average
		}
	}

	jsonData, err := json.Marshal(measuresAve)
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

// ---------------------- other function ----------------------

func formatDate(date string) (time.Time, error) {
	index := strings.LastIndex(date, "-")
	b := []byte(date)
	b[index] = ' '
	date = string(b)

	return time.Parse("2006-01-02 15:04", date)
}

func keysBetweenDate(start, end time.Time, keysMeasure []string) []string {
	y, m, d := start.Date()
	start = time.Date(y, m, d, start.Hour(), start.Minute(), 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, end.Hour(), end.Minute(), 0, 0, time.UTC)

	var res []string

	for _, key := range keysMeasure {
		var elem entities.MeasureMemKey
		json.Unmarshal([]byte(key), &elem)
		date_time, _ := formatDate(elem.Date[:len(elem.Date)-3])

		y, m, d = date_time.Date()
		date_time = time.Date(y, m, d, date_time.Hour(), date_time.Minute(), 0, 0, time.UTC)

		if (date_time.After(start) && date_time.Before(end)) || date_time.Equal(start) {
			res = append(res, key)
		}
	}
	return res
}

func average(measures []entities.MeasureValue) float32 {
	var sum float32 = 0.0
	for _, measure := range measures {
		sum += measure.Value
	}
	return sum / float32(len(measures))
}
