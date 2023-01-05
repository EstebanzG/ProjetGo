package entities

type SensorValue struct {
	AirportID     string  `json:"airport_id"`
	Date          string  `json:"date"`
	MeasureNature string  `json:"measure_nature"`
	SensorId      int     `json:"sensor_id"`
	Value         float32 `json:"value"`
}
