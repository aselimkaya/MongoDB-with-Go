package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aselimkaya/mongodb/service"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	logger := log.New(os.Stdout, "mongodb", log.LstdFlags)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//Veritabanı bağlantısı
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	//Servisi oluşturuyoruz
	personService := service.NewPersonService(client, logger)

	router := mux.NewRouter()
	router.HandleFunc("/person", personService.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", personService.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", personService.GetPersonEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}
