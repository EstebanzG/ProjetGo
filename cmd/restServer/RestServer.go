package main

import (
	"log"
	"net/http"

	"foo.org/myapp/internal/web/rest"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/data/{sensorType}", rest.GetByType)
	router.HandleFunc("/data/{sensorType}/{date1}/{date2}", rest.GetByTypeBetweenDate)
	router.HandleFunc("/average/{date}", rest.GetAverageByDay)

	router.HandleFunc("/data/{airportIATA}/{sensorType}", rest.GetByType)
	router.HandleFunc("/data/{airportIATA}/{sensorType}/{date1}/{date2}", rest.GetByTypeBetweenDate)
	router.HandleFunc("/average/{airportIATA}/{date}", rest.GetAverageByDay)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
