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
	month := int(now.Month())

	if month < 3 || month >= 12 { // Winter
		return generateByMonth(-5, 5, oldTemp)
	} else if month >= 3 || month < 6 { // Spring
		return generateByMonth(5, 15, oldTemp)
	} else if month >= 6 || month < 9 { // Summer
		return generateByMonth(15, 35, oldTemp)
	} else { // fall
		return generateByMonth(0, 15, oldTemp)
	}

}

func generateByMonth(minTemp float32, maxTemp float32, oldTemp float32) float32 {
	now := time.Now()
	hour := now.Hour()
	rand.Seed(time.Now().UnixNano())

	if oldTemp == 100.0 {
		if hour > 18 || hour < 10 {
			return float32(rand.Intn(5)) + minTemp
		} else {
			return maxTemp - float32(rand.Intn(5))
		}
	}
	var delta float32
	if hour < 6 {
		// Between -1 and 1
		delta = rand.Float32()*2 - 1
	} else if hour < 14 {
		// Between -1 and 2
		delta = rand.Float32()*3 - 1
	} else if hour < 18 {
		// Between -1 and 1
		delta = rand.Float32()*2 - 1
	} else {
		// Between -2 and 1
		delta = rand.Float32()*-3 + 1
	}
	if oldTemp+delta < minTemp || oldTemp+delta > maxTemp {
		return oldTemp
	}
	return oldTemp + delta
}
