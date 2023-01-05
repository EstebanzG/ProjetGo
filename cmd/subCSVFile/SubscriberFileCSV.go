package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Measurement struct {
	AirportID    string    `json:"airport_id"`
	Date         time.Time `json:"date"`
	MeasureNature string    `json:"measure_nature"`
	SensorID     int       `json:"sensor_id"`
	Value        int       `json:"value"`
}

func main() {
	// Parse JSON object
	var m Measurement
	err := json.Unmarshal([]byte(`{"airport_id":"NTS","date":"2023-01-04-13:46:53","measure_nature":"temperature","sensor_id":5,"value":3}`), &m)
	if err != nil {
		log.Fatal(err)
	}

	// Create CSV file
	filename := fmt.Sprintf("measurements_%s.csv", m.Date.Format("2006-01-02"))
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write measurement to CSV
	err = writer.Write([]string{m.AirportID, m.Date.Format("2006-01-02 15:04:05"), m.MeasureNature, fmt.Sprintf("%d", m.SensorID), fmt.Sprintf("%d", m.Value)})
	if err != nil {
		log.Fatal(err)
	}

	// Save CSV file every 24 hours
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()
	for {
		<-ticker.C
		writer.Flush()
		file.Close()
		file, err = os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		writer = csv.NewWriter(file)
	}
}
