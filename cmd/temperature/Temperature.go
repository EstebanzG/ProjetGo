package main

import (
	"foo.org/myapp/internal/config"
	"foo.org/myapp/internal/format"
	"math/rand"
	"time"

	"foo.org/myapp/internal/server"
)

func main() {
	pub()
}

func pub() {
	var sensorId = config.GetTemperatureSensorId()
	var airportIATA = config.GetAirportIATA()

	client := server.Connect(config.GetFullURL(), "temperature")

	for {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		temp := r1.Intn(35)
		dataToSend := format.DataToSend(sensorId, airportIATA, "temperature", float32(temp))
		client.Publish("temperature", config.GetqOs(), false, dataToSend).Wait()
		time.Sleep(10 * time.Second)
	}
}
