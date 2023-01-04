package persistence

import (
	"encoding/json"
	"fmt"
	"foo.org/myapp/internal/database"
	"foo.org/myapp/internal/entities"
	"github.com/gomodule/redigo/redis"
	"log"
	"math/rand"
	"strings"
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

func SelectAllDataForADay() ([]entities.Sensor, []entities.Sensor, []entities.Sensor) {
	conn := database.GetConnexion()
	windKeys, err := redis.Strings(conn.Do("KEYS", "wind//2023-01-04*"))
	temperatureKeys, err := redis.Strings(conn.Do("KEYS", "temperature//2023-01-04*"))
	pressureKeys, err := redis.Strings(conn.Do("KEYS", "pressure//2023-01-04*"))
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
	return GetForKeys(windKeys), GetForKeys(temperatureKeys), GetForKeys(pressureKeys)
}

func GetForKeys(keys []string) []entities.Sensor {
	conn := database.GetConnexion()
	var objects []entities.Sensor
	for _, key := range keys {
		value, _ := redis.Bytes(conn.Do("GET", key))
		objetMem := entities.SensorMem{}
		err := json.Unmarshal(value, &objetMem)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		parts := strings.Split(key, "//")
		measureNature := parts[0]
		date := parts[1]

		object := entities.Sensor{
			AirportID:     objetMem.AirportID,
			Date:          date,
			MeasureNature: measureNature,
			SensorId:      objetMem.SensorId,
			Value:         objetMem.Value,
		}
		objects = append(objects, object)
	}
	conn.Close()
	return objects
}
