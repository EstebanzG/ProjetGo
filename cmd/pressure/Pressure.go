package main

import (
	"fmt"
	"foo.org/myapp/internal/format"
	"math/rand"
	"time"

	"foo.org/myapp/internal/config"
	"foo.org/myapp/internal/server"
)

func main() {
	pub()
}

func pub() {
	var sensorId = config.GetPressureSensorId()
	var airportIATA = config.GetAirportIATA()
	var pressure float32 = 0.0

	client := server.Connect(config.GetFullURL(), "pressure")

	for {
		pressure = generate(pressure)
		fmt.Println(pressure)
		dataToSend := format.DataToSend(sensorId, airportIATA, "pressure", pressure)
		client.Publish("pressure", config.GetqOs(), false, dataToSend).Wait()
		time.Sleep(10 * time.Second)
	}
}

func generate(oldPressure float32) float32 {
	if oldPressure == 0.0 {
		rand.Seed(time.Now().UnixNano())
		// Between 990 and 1020
		return 990 + rand.Float32()*30
	}
	// Between -2 and 2
	delta := -2 + rand.Float32()*4
	if oldPressure+delta < 960 || oldPressure+delta > 1040 {
		return oldPressure
	}
	return oldPressure + delta
}
