package main

import (
	"foo.org/myapp/internal/web/rest"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	//TODO : ajouter les bornes date en paramètre
	router.HandleFunc("/data/{sensorType}", rest.GetAllSensor)
	router.HandleFunc("/moyenne/{date}", rest.GetMoyenneAllDataForADay)
	http.ListenAndServe(":8080", router)
}
