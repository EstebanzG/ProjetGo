package persistence

import (
	"foo.org/myapp/internal/entities"
	"math/rand"
)

func SelectAll() []entities.Sensor {
	numReadings := 10
	readings := make([]entities.Sensor, numReadings)
	for i := 0; i < numReadings; i++ {
		readings[i] = entities.Sensor{
			AirportID:     "NTS",
			Date:          "2023",
			MeasureNature: "NTS",
			SensorId:      i,
			Value:         rand.Float32() * 100,
		}
	}
	return readings
}
