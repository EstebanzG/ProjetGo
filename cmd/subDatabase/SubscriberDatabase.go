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
	"strconv"
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

func addToDatabase(_ mqtt.Client, message mqtt.Message) {
	objet := entities.SensorValue{}
	err := json.Unmarshal(message.Payload(), &objet)
	if err != nil {
		return
	}

	key := format.DataKeyToStore(objet.AirportID, objet.Date, objet.MeasureNature, strconv.Itoa(objet.SensorId))
	value := format.DataToStore(objet.Value)
	fmt.Println(string(key))

	conn := database.GetConnexion()
	defer database.Close(conn)
	_, err = conn.Do("SET", key, value)
	if err != nil {
		log.Fatal(err)
	}
}
