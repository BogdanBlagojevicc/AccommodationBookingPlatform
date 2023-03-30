package main

import (
	"context"
	"log"
	"os"
	"time"
	"user-service/handler"
	"user-service/repository"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	userLogger := log.New(os.Stdout, "[user-store] ", log.LstdFlags)

	userStore, err := repository.New(timeoutContext, userLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer userStore.Disconnect(timeoutContext)

	userStore.Ping()

	usersHandler := handler.NewUserHandler(logger, userStore)

	usersHandler.DatabaseName(timeoutContext)
}
