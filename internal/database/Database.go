package database

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func GetConnexion() redis.Conn {
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		log.Fatal(err)
	}

	return conn
}
