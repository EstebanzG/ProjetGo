package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"foo.org/myapp/internal/config"
	"foo.org/myapp/internal/entities"
	"foo.org/myapp/internal/server"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const csvPath = "./data/"
type Measurement struct {
	AirportID     string  `json:"airport_id"`
	Date          string  `json:"date"`
	MeasureNature string  `json:"measure_nature"`
	SensorID      int     `json:"sensor_id"`
	Value         float32 `json:"value"`
}


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

func dataFormat(message mqtt.Message) entities.MeasureValue {
	var data entities.MeasureValue
	s := message.Payload()
	fmt.Println(string(s))
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		log.Fatal(err)
	}
	return data

}
func CreateFileCSV(data entities.MeasureValue, pathname string) {
	if _, err := os.Stat(pathname); os.IsNotExist(err) {
		file, err := os.Create(pathname)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		file, err = os.OpenFile(pathname, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := file.WriteString("AirportIATA,Date,MeasureNature,SensorId,Value\n"); err != nil {
			log.Fatal(err)
		}
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}

	}
}
func WriteFileCSV(data entities.MeasureValue, pathname string) {

	file, err := os.OpenFile(pathname, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := file.WriteString(data.AirportIATA + "," + data.Date + "," + data.MeasureNature + "," + strconv.Itoa(data.SensorId) + "," + fmt.Sprintf("%.2f", data.Value) + "\n"); err != nil {
		log.Fatal(err)
	}
	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
}

func addToCSVFile(_ mqtt.Client, message mqtt.Message) {
	data := dataFormat(message)
	now := time.Now()
	pathname := csvPath + data.AirportIATA + "_" + data.MeasureNature + "_" + string(now.Format("2006-01-02")) + ".csv"
	CreateFileCSV(data, pathname)
	WriteFileCSV(data, pathname)

}
