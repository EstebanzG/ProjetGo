package main

import (
	"foo.org/myapp/internal/web/rest"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/data/{sensorType}", rest.GetAllSensor)
	http.ListenAndServe(":8080", router)
}
