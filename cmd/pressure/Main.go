package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	HTTP struct {
		Port string `yaml:"port"`
		Url string `yaml:"url"`
		Protocol string `yaml:"protocol"`
	}
}

func (c *config) getConfig() *config {
    yamlFile, err := os.ReadFile("../pub/config.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

func main() {
<<<<<<< Updated upstream:cmd/pressure/Main.go
	pub()
=======
	var c config
	var port = c.getConfig().HTTP.Port
	var url = c.getConfig().HTTP.Url
	var protocol = c.getConfig().HTTP.Protocol
	pub("test")
>>>>>>> Stashed changes:cmd/pub/Main.go
}
