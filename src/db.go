package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect() (*mongo.Client, error) {
	err := Loadenv()
	if err != nil {
		log.Fatal("Unable to load env")
		return nil, err
	}

	uri := os.Getenv("MONGODB_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to DB")
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
			panic(err)
		}
	}()

	return client, nil
}
