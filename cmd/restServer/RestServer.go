package main

import (
	"log"
	"net/http"

	"foo.org/myapp/internal/web/rest"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//TODO : ajouter les bornes date en paramètre
	router.HandleFunc("/data/{sensorType}", rest.GetBySensorType)
	router.HandleFunc("/data/{sensorType}/{date1}/{date2}", rest.GetBySensorTypeBetweenDate)
	router.HandleFunc("/average/{airportIATA}/{date}", rest.GetAverageOfAllMeasureByADay)
	router.HandleFunc("/average/{date}", rest.GetAverageOfAllMeasureByADay)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
