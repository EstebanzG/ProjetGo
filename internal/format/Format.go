package format

import (
	"encoding/json"
	"fmt"
	"foo.org/myapp/internal/entities"
	"time"
)

func FormatDataToSend(sensorId int, airportID string, measureNature string, value float32) []byte {
	object := entities.Sensor{
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
	object := entities.SensorMem{
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
