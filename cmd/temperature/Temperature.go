package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"foo.org/myapp/internal/server"
)

func pub() {
	client := server.Connect("tcp://localhost:1883", "temperature")
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
