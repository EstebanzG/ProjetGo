package main

import (
	"fmt"
)

func pub(topic string, message string) {
	fmt.Println("hello world")
	client := connect("tcp://localhost:1883", "my-client-id")
	fmt.Println("connexion ok")
	client.Publish(topic, 0, false, message).Wait()
	fmt.Println("publish ok")
}
