package main

import (
	"fmt"
	"math/rand"
	"time"

	"foo.org/myapp/internal/server"
)

func pub() {
	client := server.Connect("tcp://localhost:1883", "wind")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	wind := r1.Float32()*15 + 5.00

	for {
		//speed 1 normal or 4 acc
		wind = calculateNewWind(wind, 1.00) //strconv.Itoa(randomIndex)
		fmt.Println(wind)
		client.Publish("wind", 0, false, wind).Wait()
		time.Sleep(1 * time.Second)
	}
}

func calculateNewWind(wind float32, speed float32) float32 {
	sign := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2)

	if wind >= 22.00 {
		return wind - rand.New(rand.NewSource(time.Now().UnixNano())).Float32()*speed
	} else if wind <= 5.00 {
		return wind + rand.New(rand.NewSource(time.Now().UnixNano())).Float32()*speed
	}

	if sign == 0 {
		return wind - rand.New(rand.NewSource(time.Now().UnixNano())).Float32()*speed
	} else if sign == 1 {
		return wind + rand.New(rand.NewSource(time.Now().UnixNano())).Float32()*speed
	}

	return wind
}
