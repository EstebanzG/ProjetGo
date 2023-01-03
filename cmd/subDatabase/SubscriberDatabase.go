package main

import (
	"fmt"
	"foo.org/myapp/internal/config"
	"foo.org/myapp/internal/server"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
)

func sub() {

	client := server.Connect(config.GetFullURL(), "subscriber")
	wg := sync.WaitGroup{}
	wg.Add(1)
	//Mettre en place un client.SuscribeMultiple()
	client.Subscribe("temperature", 0, receiveTemperature)
	client.Subscribe("wind", 0, receiveWind)
	client.Subscribe("pressure", 0, receivePressure)
	wg.Wait()
}

func receiveTemperature(client mqtt.Client, message mqtt.Message) {
	fmt.Println("Temp")
	fmt.Println(string(message.Payload()))
}

func receiveWind(client mqtt.Client, message mqtt.Message) {
	fmt.Println("Wind")
	fmt.Println(string(message.Payload()))
}

func receivePressure(client mqtt.Client, message mqtt.Message) {
	fmt.Println("Pressure")
	fmt.Println(string(message.Payload()))
}
