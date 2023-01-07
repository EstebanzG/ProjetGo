package rest

import (
	"encoding/json"
	"fmt"
	"foo.org/myapp/internal/entities"
	"foo.org/myapp/internal/persistence"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"
)

// ---------------------- route function ----------------------

func GetByTypeBetweenDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	measureType := vars["measureType"]
	startDate := vars["startDate"]
	endDate := vars["endDate"]
	airportIATA := getAirportIATA(w, vars)

	matchStartDate, _ := matchDateHour(startDate)
	matchEndDate, _ := matchDateHour(endDate)
	if !matchStartDate || !matchEndDate {
		http.Error(w, "Bad Request, the date must respect the format : YYYY-MM-DD-hh:mm", http.StatusBadRequest)
		return
	}

	formatStartDate, err1 := formatDate(startDate)
	if err1 != nil {
		errInternal(w, err1)
	}
	formatEndDate, err2 := formatDate(endDate)
	if err2 != nil {
		errInternal(w, err2)
	}

	if formatStartDate.After(formatEndDate) {
		http.Error(w, "Bad Request, change date order", http.StatusBadRequest)
		return
	}

	var keysOfDate []string
	for date := formatStartDate; date.Before(formatEndDate) || date.Equal(formatEndDate); date = date.AddDate(0, 0, 1) {
		keysOfDate = append(keysOfDate, persistence.SelectKeysByDate(airportIATA, measureType, date.Format("2006-01-02"))...)
	}

	allKeys := keysBetweenDate(formatStartDate, formatEndDate, keysOfDate)
	sort.Strings(allKeys)

	if len(allKeys) == 0 {
		errNoData(w)
	}
	data := persistence.GetForKeys(allKeys)
	if len(data) == 0 {
		errNoData(w)
	}

	render(w, data)
}

func GetByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	measureType := vars["measureType"]
	airportIATA := getAirportIATA(w, vars)

	allKeys := persistence.SelectKeys(airportIATA, measureType)
	sort.Strings(allKeys)
	if len(allKeys) == 0 {
		errNoData(w)
	}
	data := persistence.GetForKeys(allKeys)
	if len(data) == 0 {
		errNoData(w)
	}
	render(w, data)
}

func GetAverageByDay(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]
	airportIATA := getAirportIATA(w, vars)

	match, _ := matchDate(date)
	if !match {
		http.Error(w, "Bad Request, the date must respect the format : YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		http.Error(w, "Bad Request, the day doesn't exist", http.StatusBadRequest)
		return
	}

	measuresAve := entities.MeasureAvg{}
	var isWind, isPressure, isTemperature bool
	measuresAve.WindAverage, isWind = averageType(airportIATA, "wind", date)
	measuresAve.PressureAverage, isPressure = averageType(airportIATA, "pressure", date)
	measuresAve.TemperatureAverage, isTemperature = averageType(airportIATA, "temperature", date)

	if !(isWind && isPressure && isTemperature) {
		errNoData(w)
	}

	render(w, measuresAve)
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
		dateTime, _ := formatDate(elem.Date[:len(elem.Date)-3])

		y, m, d = dateTime.Date()
		dateTime = time.Date(y, m, d, dateTime.Hour(), dateTime.Minute(), 0, 0, time.UTC)

		if (dateTime.After(start) && dateTime.Before(end)) || dateTime.Equal(start) {
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

func averageType(airportIATA, measureType, date string) (float32, bool) {
	measuresKeys := persistence.SelectKeysByDate(airportIATA, measureType, date)
	measures := persistence.GetForKeys(measuresKeys)
	if len(measures) != 0 {
		return average(measures), true
	}
	return 0.00, false
}

func matchAirport(airportIATA string) (bool, error) {
	return regexp.MatchString("^[A-Z]{3}$", airportIATA)
}

func matchDate(date string) (bool, error) {
	return regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}$", date)
}

func matchDateHour(date string) (bool, error) {
	return regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2}:[0-9]{2}$", date)
}

func getAirportIATA(w http.ResponseWriter, vars map[string]string) string {
	airportIATA, exists := vars["airportIATA"]
	if exists {
		match, _ := matchAirport(airportIATA)
		if !match {
			errAirportIATA(w)
		}
	} else {
		airportIATA = "*"
	}
	return airportIATA
}

// ---------------------- error and render function ----------------------
func render(w http.ResponseWriter, data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		errInternal(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func errNoData(w http.ResponseWriter) {
	http.Error(w, "No data available for this day", http.StatusNoContent)
	return
}

func errAirportIATA(w http.ResponseWriter) {
	http.Error(w, "Bad Request, the airport IATA code is invalid", http.StatusBadRequest)
	return
}

func errInternal(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}
