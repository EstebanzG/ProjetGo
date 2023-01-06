package main

import (
	"net/http"

	"foo.org/myapp/internal/web/rest"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//TODO : ajouter les bornes date en paramètre
	router.HandleFunc("/data/{sensorType}", rest.GetAllSensor)
	//TODO : passer la date en paramètre
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/moyenne/{date}", rest.GetMoyenneAllDataForADay)
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}
