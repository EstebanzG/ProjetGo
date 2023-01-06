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
	router.HandleFunc("/data/{measureType}", rest.GetByType)
	router.HandleFunc("/data/{measureType}/{startDate}/{endDate}", rest.GetByTypeBetweenDate)
	router.HandleFunc("/average/{date}", rest.GetAverageByDay)

	router.HandleFunc("/data/{airportIATA}/{measureType}", rest.GetByType)
	router.HandleFunc("/data/{airportIATA}/{measureType}/{startDate}/{endDate}", rest.GetByTypeBetweenDate)
	router.HandleFunc("/average/{airportIATA}/{date}", rest.GetAverageByDay)
	// NE JAMAIS SUPPRIMER CES LIGNES SINON ÇA MARCHE PLUS ! OK ?!
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	err := http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
	if err != nil {
		log.Fatal(err)
		return
	}
}
