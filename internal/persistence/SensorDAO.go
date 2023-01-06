package persistence

import (
	"encoding/json"
	"foo.org/myapp/internal/database"
	"foo.org/myapp/internal/entities"
	"foo.org/myapp/internal/format"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
)

func SelectKeys(airportIATA string, measureType string) []string {
	conn := database.GetConnexion()
	defer database.Close(conn)

	keyFormat := format.DataKeyToStore(airportIATA, "*", measureType, "*")
	keys, err := redis.Strings(conn.Do("KEYS", keyFormat))
	if err != nil {
		log.Fatal(err)
	}
	return keys
}

func SelectKeysByDate(airportIATA string, measureType string, date string) []string {
	conn := database.GetConnexion()
	defer database.Close(conn)

	keyFormat := format.DataKeyToStore(airportIATA, date+"-*", measureType, "*")
	keys, err := redis.Strings(conn.Do("KEYS", keyFormat))
	if err != nil {
		log.Fatal(err)
	}
	return keys
}

func GetForKeys(keys []string) []entities.MeasureValue {
	conn := database.GetConnexion()
	defer database.Close(conn)

	var objects []entities.MeasureValue
	for _, key := range keys {
		value, _ := redis.Bytes(conn.Do("GET", key))
		objectMem := entities.MeasureMem{}
		err := json.Unmarshal(value, &objectMem)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		objectMemKey := entities.MeasureMemKey{}
		err = json.Unmarshal([]byte(key), &objectMemKey)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		sensorId, _ := strconv.Atoi(objectMemKey.SensorId)
		object := entities.MeasureValue{
			AirportIATA:   objectMemKey.AirportIATA,
			Date:          objectMemKey.Date,
			MeasureNature: objectMemKey.MeasureNature,
			SensorId:      sensorId,
			Value:         objectMem.Value,
		}

		objects = append(objects, object)
	}
	return objects
}
