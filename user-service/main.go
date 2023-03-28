package main

/*
import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"user-service/handler"
	"user-service/repository"
	"user-service/service"
)

func initDB() *mongo.Client {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	mongoUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoUri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")

	return client
}

func initRepo(client *mongo.Client) *repository.UserRepository {
	return &repository.UserRepository{Client: client}
}

func initService(repo *repository.UserRepository) *service.UserService {
	return &service.UserService{Repo: repo}
}

func initHandler(service *service.UserService) *handler.UserHandler {
	return &handler.UserHandler{Service: service}
}

func handleFunc(handler *handler.UserHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.Hello).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}

func main() {

	client := initDB()
	repository := initRepo(client)
	service := initService(repository)
	handler := initHandler(service)
	handleFunc(handler)
}
*/
