package entities

type SensorMem struct {
	AirportID string  `json:"airport_id"`
	SensorId  int     `json:"sensor_id"`
	Value     float32 `json:"value"`
}
