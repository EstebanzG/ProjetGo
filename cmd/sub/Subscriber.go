package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"sync"
	"time"
)

func sub(topic string) {
	client := connect("tcp://localhost:1883", "subscriber")
	wg := sync.WaitGroup{}
	wg.Add(1)
	client.Subscribe(topic, 0, myfunction)
	wg.Wait()
}

func myfunction(client mqtt.Client, message mqtt.Message) {
	fmt.Println(string(message.Payload()))
}

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	return opts
}

func connect(brokerURI string, clientId string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ")...")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}
