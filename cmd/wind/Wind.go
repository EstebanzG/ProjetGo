package main

import (
	"fmt"
	"foo.org/myapp/internal/config"
	"math/rand"
	"time"

	"foo.org/myapp/internal/server"
)

func pub() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	wind := r1.Float32()*15 + 5.00
	client := server.Connect(config.GetFullURL(), "wind")

	for {
		wind = calculateNewWind(wind, 1.00) //strconv.Itoa(randomIndex)
		fmt.Println(wind)
		windString := fmt.Sprintf("%f", wind)
		client.Publish("wind", 0, false, windString).Wait()
		time.Sleep(10 * time.Second)
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
