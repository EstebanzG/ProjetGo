package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"foo.org/myapp/internal/config"
	"foo.org/myapp/internal/server"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Measurement struct {
	AirportID     string    `json:"airport_id"`
	Date          time.Time `json:"date"`
	MeasureNature string    `json:"measure_nature"`
	SensorID      int       `json:"sensor_id"`
	Value         int       `json:"value"`
}

var csvFile *os.File

func main() {
	subCSV()
}

func subCSV() {
	client := server.Connect(config.GetFullURL(), "subscriberCSV")
	wg := sync.WaitGroup{}
	wg.Add(1)
	client.Subscribe("temperature", 0, addToCSVFile)
	client.Subscribe("wind", 0, addToCSVFile)
	client.Subscribe("pressure", 0, addToCSVFile)
	wg.Wait()
}

func addToCSVFile(_ mqtt.Client, message mqtt.Message) {

	now := time.Now()
	csvFile, err := os.Create(now.Format("2006-01-02-15:04:05") + ".csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	
	// print the message.payload
	fmt.Println(string(message.Payload()["airport_id"]))
	// Create a new CSV writer
	csvWriter := csv.NewWriter(csvFile)
	err = csvWriter.Write(string{message.Payload().airport_id, data["date"], data["measure_nature"], data["sensor_id"]})
	if err != nil {
		log.Println(err)
	}

	defer csvWriter.Flush()
	// Save the file every 24 hours
	for {
		time.Sleep(60 * time.Second)

		// Close the current file and create a new one with the updated date
		csvFile.Close()
		now = time.Now()
		csvFile, err = os.Create(now.Format("2006-01-02-15:04:05") + ".csv")
		if err != nil {
			log.Fatal(err)
		}
		csvWriter = csv.NewWriter(csvFile)
	}
}
