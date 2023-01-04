package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type config struct {
	HTTP struct {
		Port     string
		Url      string `yaml:"url"`
		Protocol string `yaml:"protocol"`
	} `yaml:"http"`
	Info struct {
		WindSensorId        int    `yaml:"wind-sensor-id"`
		TemperatureSensorId int    `yaml:"temperature-sensor-id"`
		PressureSensorId    int    `yaml:"pressure-sensor-id"`
		AirportId           string `yaml:"airport-id"`
	} `yaml:"info"`
}

var c *config

func init() {
	c = new(config)
	c.getConfig()
}

func (conf *config) getConfig() *config {
	yamlFile, err := os.ReadFile("internal/config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return conf
}

func getPort() string {
	return c.getConfig().HTTP.Port
}

func getURL() string {
	return c.getConfig().HTTP.Url
}

func getProtocol() string {
	return c.getConfig().HTTP.Protocol
}

func GetFullURL() string {
	fullURL := getProtocol() + "://" + getURL() + ":" + getPort()
	return fullURL
}

func GetWindSensorId() int {
	return c.getConfig().Info.WindSensorId
}

func GetTemperatureSensorId() int {
	return c.getConfig().Info.TemperatureSensorId
}

func GetPressureSensorId() int {
	return c.getConfig().Info.PressureSensorId
}

func GetAirportId() string {
	return c.getConfig().Info.AirportId
}
