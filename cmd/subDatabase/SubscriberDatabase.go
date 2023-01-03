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
	client.Subscribe("temperature", 0, received)
	client.Subscribe("wind", 0, received)
	client.Subscribe("pressure", 0, received)
	wg.Wait()
}

func received(client mqtt.Client, message mqtt.Message) {
	fmt.Println(string(message.Payload()))
}
