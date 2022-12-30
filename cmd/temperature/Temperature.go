package main

import (
	"fmt"
	"foo.org/myapp/internal/config"
	"math/rand"
	"strconv"
	"time"

	"foo.org/myapp/internal/server"
)

func pub() {
	client := server.Connect(config.GetFullURL(), "temperature")
	for {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		randomIndex := r1.Intn(35)
		temp := strconv.Itoa(randomIndex) + "°C"
		fmt.Println(temp)
		client.Publish("temperature", 0, false, temp).Wait()
		time.Sleep(10 * time.Second)
	}
}
