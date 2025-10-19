package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoResult struct {
	ID    string `json:"_id"`
	Short string `json:"short"`
	Long  string `json:"long"`
}

type URLRequestType struct {
	Url string `json:"url"`
}

type URLResponseType struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Url     string `json:"url"`
}

type Handlers struct {
	Client *mongo.Client
}

func (h *Handlers) redirectURL(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	db := h.Client.Database("url-shortner")
	collection := db.Collection("links")

	filter := bson.M{
		"short": id,
	}

	var findResult MongoResult

	err := collection.FindOne(context.TODO(), filter).Decode(&findResult)
	if err != nil {
		fmt.Println("Error fetching")
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println("Redirecting to:", findResult.Long)
	http.Redirect(w, r, findResult.Long, http.StatusMovedPermanently)
}

func (h *Handlers) generateURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "86400")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var result URLRequestType
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	db := h.Client.Database("url-shortner")
	collection := db.Collection("links")

	id, err := gonanoid.New(6)
	if err != nil {
		http.Error(w, "Unable to generate a nanoid", http.StatusInternalServerError)
		log.Fatalf("Unable to generate id: %v", err)
		return
	}

	newLink := bson.D{
		{Key: "short", Value: id},
		{Key: "long", Value: result.Url},
	}

	res, err := collection.InsertOne(context.TODO(), newLink)
	if err != nil {
		log.Fatal("Insert failed", err)
	}

	fmt.Println("Document inserted", res.InsertedID)

	w.WriteHeader(http.StatusOK)
	response := URLResponseType{
		Success: true,
		Message: "Generated URL",
		Url:     "http://localhost:8080/" + id,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
