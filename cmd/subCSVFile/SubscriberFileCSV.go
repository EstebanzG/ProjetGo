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
	AirportIATA   string    `json:"airport_iata"`
	Date          time.Time `json:"date"`
	MeasureNature string    `json:"measure_nature"`
	SensorID      int       `json:"sensor_id"`
	Value         int       `json:"value"`
}

var csvFile *os.File

func main() {
	// Parse JSON object
	var m Measurement
	err := json.Unmarshal([]byte(`{"airport_iata":"NTS","date":"2023-01-04-13:46:53","measure_nature":"temperature","sensor_id":5,"value":3}`), &m)
	if err != nil {
		log.Fatal(err)
	}

func addToCSVFile(_ mqtt.Client, message mqtt.Message) {

	now := time.Now()
	csvFile, err := os.Create(now.Format("2006-01-02-15:04:05") + ".csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write measurement to CSV
	err = writer.Write([]string{m.AirportIATA, m.Date.Format("2006-01-02 15:04:05"), m.MeasureNature, fmt.Sprintf("%d", m.SensorID), fmt.Sprintf("%d", m.Value)})
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
