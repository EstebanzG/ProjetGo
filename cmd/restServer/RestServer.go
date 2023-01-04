package main

import (
	"foo.org/myapp/internal/web/rest"
	"net/http"
)

func main() {
	http.HandleFunc("/temp", rest.GetAllSensor)
	http.ListenAndServe(":8080", nil)
}
