package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func pub() {
	client := connect("tcp://localhost:1883", "wind")
	for {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		randomIndex := r1.Intn(35)
		temp := strconv.Itoa(randomIndex) + "°C"
		fmt.Println(temp)
		client.Publish("wind", 0, false, temp).Wait()
		time.Sleep(3 * time.Second)
	}
}

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	return opts
}

func connect(brokerURI string, clientId string) mqtt.Client {

	fmt.Println("Connexion ok (" + brokerURI + ", " + clientId + ")")
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
