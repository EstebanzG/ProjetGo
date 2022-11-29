package server

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

<<<<<<< Updated upstream:internal/server/Server.go
=======
func sub(topic string) {
	fmt.Println("Subscriber OK")
	client := connect("tcp://localhost:1883", "subscriber")
	wg := sync.WaitGroup{}
	wg.Add(1)
	client.Subscribe(topic, 0, myfunction)
	wg.Wait()
}

func myfunction(client mqtt.Client, message mqtt.Message) {
	temp := message.Payload()
	fmt.Println("Reception new temp : " + string(temp))
}

>>>>>>> Stashed changes:cmd/sub/Subscriber.go
func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	return opts
}

<<<<<<< Updated upstream:internal/server/Server.go
func Connect(brokerURI string, clientId string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ")...")
=======
func connect(brokerURI string, clientId string) mqtt.Client {
	fmt.Println("Connexion sub ok (" + brokerURI + ", " + clientId + ")...")
>>>>>>> Stashed changes:cmd/sub/Subscriber.go
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
