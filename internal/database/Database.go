package database

import (
	"fmt"
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

func Close(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Println(err)
	}
}
