package entities

type SensorMemKey struct {
	AirportID     string `json:"airport_id"`
	Date          string `json:"date"`
	MeasureNature string `json:"measure_nature"`
	SensorId      string `json:"sensor_id"`
}
