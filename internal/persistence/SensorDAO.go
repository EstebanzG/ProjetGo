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

func SelectByType(sensorType string) []entities.MeasureValue {
	conn := database.GetConnexion()
	defer database.Close(conn)
	keysFormat := format.DataKeyToStore("*", "*", sensorType, "*")
	keys, err := redis.Strings(conn.Do("KEYS", keysFormat))
	if err != nil {
		log.Fatal(err)
	}
	return GetForKeys(keys)
}

func SelectKeysByType(sensorType string) []string {
	conn := database.GetConnexion()
	defer database.Close(conn)

	keyFormat := format.DataKeyToStore("*", "*", sensorType, "*")
	keys, err := redis.Strings(conn.Do("KEYS", keyFormat))
	if err != nil {
		log.Fatal(err)
	}
	return keys
}

func SelectAllDataForADay(airportIATA string, date string) map[string][]entities.MeasureValue {
	conn := database.GetConnexion()
	defer database.Close(conn)

	measuresNatures := []string{"wind", "temperature", "pressure"}
	allMeasures := make(map[string][]entities.MeasureValue)

	for _, measureNature := range measuresNatures {
		keyFormat := format.DataKeyToStore(airportIATA, date+"*", measureNature, "*")
		keys, err := redis.Strings(conn.Do("KEYS", keyFormat))
		if err != nil {
			log.Fatal(err)
		}
		if len(keys) != 0 {
			allMeasures[measureNature] = GetForKeys(keys)
		}
	}
	return allMeasures
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
