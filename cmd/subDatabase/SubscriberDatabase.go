package main

import (
	"fmt"
	"foo.org/myapp/internal/server"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
)

func sub() {
	client := server.Connect("tcp://localhost:1883", "subscriber")
	wg := sync.WaitGroup{}
	wg.Add(1)
	client.Subscribe("temperature", 0, myfunction)
	client.Subscribe("wind", 0, myfunction)
	client.Subscribe("pressure", 0, myfunction)

	wg.Wait()
}

func myfunction(client mqtt.Client, message mqtt.Message) {
	fmt.Println(string(message.Payload()))
}
