package entities

type SensorAvg struct {
	WindAverage        float32 `json:"wind-average"`
	TemperatureAverage float32 `json:"temperature_average"`
	PressureAverage    float32 `json:"pressure_average"`
}
