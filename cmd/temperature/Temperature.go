package main

import (
	"fmt"
	"foo.org/myapp/internal/config"
	"foo.org/myapp/internal/format"
	"math/rand"
	"time"

	"foo.org/myapp/internal/server"
)

var idCapteur = config.GetTemperatureSensorId()
var airportId = config.GetAirportId()

func pub() {
	client := server.Connect(config.GetFullURL(), "temperature")

	for {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		temp := r1.Intn(35)
		fmt.Println(temp)
		dataToSend := format.FormatData(idCapteur, airportId, "temperature", float32(temp))
		client.Publish("temperature", 0, false, dataToSend).Wait()
		time.Sleep(10 * time.Second)
	}
}
