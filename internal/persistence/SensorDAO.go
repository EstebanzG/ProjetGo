package persistence

import (
	"foo.org/myapp/internal/database"
	"foo.org/myapp/internal/entities"
	"log"
	"math/rand"
)

func SelectDataSensorFromTo(sensorType string) []entities.Sensor {
	conn := database.GetConnexion()
	res, err := conn.Do("KEYS " + sensorType + "//*")
	if err != nil {
		log.Fatal(err)
	}

	print(res)
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

	conn.Close()
	//numReadings := 10
	//readings := make([]entities.Sensor, numReadings)
	//for i := 0; i < numReadings; i++ {
	//	readings[i] = entities.Sensor{
	//		AirportID:     "NTS",
	//		Date:          "2023",
	//		MeasureNature: "NTS",
	//		SensorId:      i,
	//		Value:         rand.Float32() * 100,
	//	}
	//}
	return readings
}
