package main

import (
	"log"
	"net/http"

	"foo.org/myapp/internal/web/rest"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//TODO : passer la date en paramètre
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
	router.HandleFunc("/data/{sensorType}", rest.GetBySensorType)
	router.HandleFunc("/data/{sensorType}/{date1}/{date2}", rest.GetDataSensorBetweenDate)
	router.HandleFunc("/average/{airportIATA}/{date}", rest.GetAverageOfAllMeasureByADay)
	router.HandleFunc("/average/{date}", rest.GetAverageOfAllMeasureByADay)
	err := http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
	if err != nil {
		log.Fatal(err)
		return
	}
}
