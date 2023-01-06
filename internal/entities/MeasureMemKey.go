package entities

type MeasureMemKey struct {
	AirportIATA   string `json:"airport_iata"`
	Date          string `json:"date"`
	MeasureNature string `json:"measure_nature"`
	SensorId      string `json:"sensor_id"`
}
