package main

import (
	"log"
)

func main() {
	client, err := Connect()
	if err != nil {
		log.Fatalf("Error connecting to database %v", err)
	}
}
