package main

import (
	"encoding/json"
	"fmt"
	"foo.org/myapp/internal/config"
	"foo.org/myapp/internal/database"
	"foo.org/myapp/internal/entities"
	"foo.org/myapp/internal/format"
	"foo.org/myapp/internal/server"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"sync"
)

func main() {
	sub()
}

func sub() {
	client := server.Connect(config.GetFullURL(), "subscriber")
	wg := sync.WaitGroup{}
	wg.Add(1)
	client.Subscribe("temperature", 0, addToDatabase)
	client.Subscribe("wind", 0, addToDatabase)
	client.Subscribe("pressure", 0, addToDatabase)
	wg.Wait()
}

func addToDatabase(client mqtt.Client, message mqtt.Message) {
	objet := entities.Sensor{}
	err := json.Unmarshal(message.Payload(), &objet)
	if err != nil {
		return
	}

	cle := objet.MeasureNature + "//" + objet.Date
	value := format.FormatDataToStore(objet.SensorId, objet.AirportID, objet.Value)
	fmt.Println(cle)

	conn := database.GetConnexion()
	_, err = conn.Do("SET", cle, value)
	if err != nil {
		log.Fatal(err)
	}

	conn.Close()
}
