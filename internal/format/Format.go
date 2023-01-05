package format

import (
	"encoding/json"
	"fmt"
	"foo.org/myapp/internal/entities"
	"time"
)

func DataToSend(sensorId int, airportIATA string, measureNature string, value float32) []byte {
	object := entities.MeasureValue{
		AirportIATA:   airportIATA,
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

func DataKeyToStore(airportIATA string, date string, measureNature string, sensorId string) []byte {
	object := entities.MeasureMemKey{
		AirportIATA:   airportIATA,
		Date:          date,
		MeasureNature: measureNature,
		SensorId:      sensorId,
	}

	jsonObject, err := json.Marshal(object)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return nil
	}
	return jsonObject
}

func DataToStore(value float32) []byte {
	object := entities.MeasureMem{
		Value: value,
	}

	jsonObject, err := json.Marshal(object)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return nil
	}
	return jsonObject
}
