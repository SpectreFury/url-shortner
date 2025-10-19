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

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Working it"))
	})

	port := os.Getenv("PORT")

	fmt.Println("Listening on PORT: ", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Error starting the server %v", err)
		panic(err)
	}

}
