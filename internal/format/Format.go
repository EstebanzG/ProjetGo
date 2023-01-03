package format

import (
	"encoding/json"
	"fmt"
	"time"
)

func FormatData(sensorId int, airportID int, measureNature string, value float32) []byte {
	data := map[string]interface{}{
		"sensorID":      sensorId,
		"airportID":     airportID,
		"measureNature": measureNature,
		"value":         value,
		"date":          time.Now().Format("2006-01-02-15:04:05"),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return nil
	}
	fmt.Printf("json data: %s\n", jsonData)
	return jsonData
}
