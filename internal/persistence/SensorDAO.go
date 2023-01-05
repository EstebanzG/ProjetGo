package persistence

import (
	"encoding/json"
	"fmt"
	"foo.org/myapp/internal/database"
	"foo.org/myapp/internal/entities"
	"foo.org/myapp/internal/format"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
	"time"
)

func SelectByType(sensorType string) []entities.Sensor {
	conn := database.GetConnexion()
	keysFormat := format.DataKeyToStore("*", "*", sensorType, "*")
	keys, err := redis.Strings(conn.Do("KEYS", keysFormat))
	if err != nil {
		log.Fatal(err)
	}
	return GetForKeys(keys)
}

func SelectAllSensorTypeDateHour(sensorType string, allDates []time.Time) []entities.Sensor {
	conn := database.GetConnexion()
	var res []entities.Sensor

	for _, d := range allDates {
		//delete minutes and seconds
		s := []rune(d.Format("2006-01-02 15:04:05"))
		sCut := string(s[:len(s)-6])

		cle := format.DataKeyToStore("*", sCut+"*", sensorType, "*")
		keys, err := redis.Strings(conn.Do("KEYS", cle))
		if err != nil {
			log.Fatal(err)
		}
		if len(keys) != 0 {
			res = append(res, GetForKeys(keys)...)
		}
	}
	conn.Close()
	return res
}

func SelectAllDataForADay(date string) map[string][]entities.Sensor {
	conn := database.GetConnexion()
	measuresNatures := []string{"wind", "temperature", "pressure"}
	allMeasures := make(map[string][]entities.Sensor)

	for _, measureNature := range measuresNatures {
		cle := format.DataKeyToStore("*", date+"*", measureNature, "*")
		keys, err := redis.Strings(conn.Do("KEYS", cle))
		if err != nil {
			log.Fatal(err)
		}
		if len(keys) != 0 {
			allMeasures[measureNature] = GetForKeys(keys)
		}
	}
	conn.Close()
	return allMeasures
}

func GetForKeys(keys []string) []entities.Sensor {
	conn := database.GetConnexion()
	var objects []entities.Sensor
	for _, key := range keys {
		value, _ := redis.Bytes(conn.Do("GET", key))
		objectMem := entities.SensorMem{}
		err := json.Unmarshal(value, &objectMem)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		objectMemKey := entities.SensorMemKey{}
		err = json.Unmarshal([]byte(key), &objectMemKey)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		sensorId, _ := strconv.Atoi(objectMemKey.SensorId)

		object := entities.Sensor{
			AirportID:     objectMemKey.AirportID,
			Date:          objectMemKey.Date,
			MeasureNature: objectMemKey.MeasureNature,
			SensorId:      sensorId,
			Value:         objectMem.Value,
		}
		objects = append(objects, object)
	}
	conn.Close()
	return objects
}
