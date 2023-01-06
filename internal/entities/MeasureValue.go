package entities

type MeasureValue struct {
	AirportIATA   string  `json:"airport_iata"`
	Date          string  `json:"date"`
	MeasureNature string  `json:"measure_nature"`
	SensorId      int     `json:"sensor_id"`
	Value         float32 `json:"value"`
}
