package persistence

import (
	"fmt"
	"foo.org/myapp/internal/database"
	"foo.org/myapp/internal/entities"
	"github.com/gomodule/redigo/redis"
	"log"
	"math/rand"
)

func SelectDataSensorFromTo(sensorType string) []entities.Sensor {
	conn := database.GetConnexion()
	keys, err := redis.Strings(conn.Do("KEYS", sensorType+"//*"))
	if err != nil {
		log.Fatal(err)
	}
	for _, key := range keys {
		fmt.Println(key)
	}
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
