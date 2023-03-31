package main

import (
	"context"
	"flight-service/handler"
	"flight-service/repository"
	"log"
	"os"
	"time"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	flightLogger := log.New(os.Stdout, "[user-store] ", log.LstdFlags)

	flightStore, err := repository.New(timeoutContext, flightLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer flightStore.Disconnect(timeoutContext)

	flightStore.Ping()

	flightsHandler := handler.NewFlightHandler(logger, flightStore)

	flightsHandler.DatabaseName(timeoutContext)
}
