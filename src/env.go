package main

import (
	"log"

	"github.com/joho/godotenv"
)

func Loadenv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file")
		return err
	}

	return nil
}
