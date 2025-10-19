package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	client, err := Connect()
	if err != nil {
		log.Fatalf("Error connecting to database %v", err)
	}
	defer Disconnect(client)

	mux := http.NewServeMux()

	handler := Handlers{
		Client: client,
	}

	mux.HandleFunc("/{id}", handler.redirectURL)
	mux.HandleFunc("/", handler.generateURL)

	port := os.Getenv("PORT")

	fmt.Println("Listening on PORT:", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Error starting the server %v", err)
		panic(err)
	}
}
