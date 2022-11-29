package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
)

func sub(topic string) {
	fmt.Println("hello world")
	client := connect("tcp://localhost:1883", "my-client-id")
	fmt.Println("connexion ok")

	wg := sync.WaitGroup{}
	wg.Add(1)

	client.Subscribe(topic, 0, myfunction)
	fmt.Println("subscribe ok")

	wg.Wait()
}

func myfunction(client mqtt.Client, message mqtt.Message) {
	strMessage := string(message.Payload())
	fmt.Println(strMessage)
}
