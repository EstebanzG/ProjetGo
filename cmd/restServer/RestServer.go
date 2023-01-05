package main

import (
	"net/http"

	"foo.org/myapp/internal/web/rest"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//TODO : ajouter les bornes date en paramètre
	router.HandleFunc("/data/{sensorType}", rest.GetAllSensor)
	//TODO : passer la date en paramètre
	router.HandleFunc("/moyenne/{date}", rest.GetMoyenneAllDataForADay)
	http.ListenAndServe(":8080", router)
}
