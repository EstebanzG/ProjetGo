package format

import (
	"encoding/json"
	"fmt"
	"time"
)

type DataSend struct {
	AirportID     string  `json:"airport_id"`
	Date          string  `json:"date"`
	MeasureNature string  `json:"measure_nature"`
	SensorId      int     `json:"sensor_id"`
	Value         float32 `json:"value"`
}

type DataStore struct {
	AirportID string  `json:"airport_id"`
	SensorId  int     `json:"sensor_id"`
	Value     float32 `json:"value"`
}

func FormatDataToSend(sensorId int, airportID string, measureNature string, value float32) []byte {
	object := DataSend{
		AirportID:     airportID,
		Date:          time.Now().Format("2006-01-02-15:04:05"),
		MeasureNature: measureNature,
		SensorId:      sensorId,
		Value:         value,
	}

	jsonObject, err := json.Marshal(object)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return nil
	}
	return jsonObject
}

func FormatDataToStore(sensorId int, airportID string, value float32) []byte {
	object := DataStore{
		AirportID: airportID,
		SensorId:  sensorId,
		Value:     value,
	}

	jsonObject, err := json.Marshal(object)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return nil
	}
	return jsonObject
}
