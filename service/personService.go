package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/aselimkaya/mongodb/model"
)

type personService struct {
	client *mongo.Client
	logger *log.Logger
}

func NewPersonService(client *mongo.Client, logger *log.Logger) *personService {
	return &personService{client, logger}
}

func (service *personService) CreatePersonEndpoint(responseWriter http.ResponseWriter, request *http.Request) {
	service.logger.Println("Yeni kişi oluşturma isteği geldi!")

	responseWriter.Header().Add("content-type", "application/json")

	person := model.Person{}

	json.NewDecoder(request.Body).Decode(&person)

	collection := service.client.Database("persondb").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//DB'ye yeni oluşturduğumuz kişiyi yazıyoruz
	result, err := collection.InsertOne(ctx, person)

	if err != nil {
		http.Error(responseWriter, "Veri yazılırken bir hata ile karşılaşıldı!: "+err.Error(), http.StatusBadRequest)
		return
	}

	service.logger.Println("Kişi oluşturuldu!")
	service.logger.Println(person)

	json.NewEncoder(responseWriter).Encode(result)
}

func (service *personService) GetPeopleEndpoint(responseWriter http.ResponseWriter, request *http.Request) {
	service.logger.Println("Tüm kişileri çağırma isteği geldi!")

	responseWriter.Header().Add("content-type", "application/json")

	people := model.GetEmptyPeopleSlice()

	collection := service.client.Database("persondb").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		http.Error(responseWriter, "Veriler getirilirken bir hata ile karşılaşıldı!", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(ctx)

	//Bütün kişilere ulaşmak için cursor yardımıyla gezmek gerekiyor
	for cursor.Next(ctx) {
		person := model.Person{}
		cursor.Decode(&person)
		people = append(people, person)
	}

	if err := cursor.Err(); err != nil {
		http.Error(responseWriter, "Veriler işlenirken bir hata ile karşılaşıldı!", http.StatusInternalServerError)
		return
	}

	service.logger.Println("Kişiler getirildi!")
	service.logger.Println(people)
	json.NewEncoder(responseWriter).Encode(people)
}

func (service *personService) GetPersonEndpoint(responseWriter http.ResponseWriter, request *http.Request) {
	service.logger.Println("Spesifik bir kişiyi çağırma isteği geldi!")

	responseWriter.Header().Add("content-type", "application/json")

	params := mux.Vars(request)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	person := model.Person{}

	collection := service.client.Database("persondb").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//Bu sefer bson yerine Person{ID: id} kullandık
	err := collection.FindOne(ctx, model.Person{ID: id}).Decode(&person)

	if err != nil {
		http.Error(responseWriter, "Kişi bulunurken bir hata ile karşılaşıldı!", http.StatusInternalServerError)
		return
	}

	service.logger.Println("Kişi getirildi!")
	service.logger.Println(person)
	json.NewEncoder(responseWriter).Encode(person)
}
