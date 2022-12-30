package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"foo.org/myapp/internal/config"
	"foo.org/myapp/internal/server"
)

func pub() {
	client := server.Connect(config.GetFullURL(), "pressure")

	for {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		randomIndex := r1.Intn(35)
		temp := strconv.Itoa(randomIndex) + "°C"
		fmt.Println(temp)
		client.Publish("pressure", 0, false, temp).Wait()
		time.Sleep(10 * time.Second)
	}
}
