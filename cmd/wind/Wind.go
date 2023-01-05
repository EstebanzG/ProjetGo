package main

import (
	"fmt"
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
	var sensorId = config.GetWindSensorId()
	var airportIATA = config.GetAirportIATA()
	var wind float32 = -1.0

	client := server.Connect(config.GetFullURL(), "wind")

	for {
		wind = generate(wind)
		fmt.Println(wind)
		dataToSend := format.DataToSend(sensorId, airportIATA, "wind", float32(wind))
		client.Publish("wind", config.GetqOs(), false, dataToSend).Wait()
		time.Sleep(10 * time.Second)
	}
}

func generate(oldWind float32) float32 {
	if oldWind == -1.0 {
		rand.Seed(time.Now().UnixNano())
		// Between 10 and 25
		return float32(10 + rand.Intn(15))
	}
	// Between -1 and +1
	delta := float32(-1 + rand.Intn(2))
	if oldWind+delta < 5 || oldWind+delta > 130 {
		return oldWind
	}
	return oldWind + delta
}
