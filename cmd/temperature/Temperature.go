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
	var sensorId = config.GetTemperatureSensorId()
	var airportIATA = config.GetAirportIATA()
	var temp float32 = 100.0

	client := server.Connect(config.GetFullURL(), "temperature")

	for {
		temp = generate(temp)
		fmt.Println(temp)
		dataToSend := format.DataToSend(sensorId, airportIATA, "temperature", temp)
		client.Publish("temperature", config.GetqOs(), false, dataToSend).Wait()
		time.Sleep(10 * time.Second)
	}
}

func generate(oldTemp float32) float32 {
	now := time.Now()
	hour := now.Hour()
	if oldTemp == 100.0 {
		rand.Seed(time.Now().UnixNano())
		if hour > 18 || hour < 10 {
			// Between - 3 and 10
			return rand.Float32()*10 - 3
		} else {
			// Between 5 and 30
			return (rand.Float32()*5 + 1) * 5
		}
	}
	var delta float32
	if hour < 6 {
		// Between -1 and 0,5
		delta = rand.Float32()*1.5 - 1

	} else if hour < 14 {
		// Between -0,5 and 2
		delta = rand.Float32()*2.5 - 0.5
	} else if hour < 17 {
		// Between 1,5 and 0,5
		delta = rand.Float32() + 0.5
	} else {
		// Between -2 and +0,5
		delta = rand.Float32()*-2.5 + 0.5
	}
	if oldTemp+delta < -5 || oldTemp+delta > 30 {
		return oldTemp
	}
	return oldTemp + delta
}
