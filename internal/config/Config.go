package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	HTTP struct {
		Port     string `yaml:"port"`
		Url      string `yaml:"url"`
		Protocol string `yaml:"protocol"`
	}
}

func (c *config) getConfig() *config {
	yamlFile, err := os.ReadFile("../../internal/config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func getPort() string {
	var c config
	return c.getConfig().HTTP.Port
}

func getURL() string {
	var c config
	return c.getConfig().HTTP.Url
}

func getProtocol() string {
	var c config
	return c.getConfig().HTTP.Protocol
}

func GetFullURL() string {
	fullURL := getProtocol() + "://" + getURL() + ":" + getPort()
	return fullURL
}
