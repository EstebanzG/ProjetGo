package main

import (
	"fmt"
	"foo.org/myapp/internal/format"
	"math/rand"
	"time"

	"foo.org/myapp/internal/config"
	"foo.org/myapp/internal/server"
)

var idCapteur = config.GetPressureSensorId()
var airportId = config.GetAirportId()

func pub() {
	client := server.Connect(config.GetFullURL(), "pressure")

	for {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		press := r1.Intn(35)
		fmt.Println(press)
		dataToSend := format.FormatData(idCapteur, airportId, "pressure", float32(press))
		client.Publish("pressure", 0, false, dataToSend).Wait()
		time.Sleep(10 * time.Second)
	}
}
