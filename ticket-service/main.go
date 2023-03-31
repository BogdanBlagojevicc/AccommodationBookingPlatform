package main

import (
	"context"
	"log"
	"os"
	"ticket-service/handler"
	"ticket-service/repository"
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
	ticketLogger := log.New(os.Stdout, "[user-store] ", log.LstdFlags)

	ticketStore, err := repository.New(timeoutContext, ticketLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer ticketStore.Disconnect(timeoutContext)

	ticketStore.Ping()

	ticketsHandler := handler.NewTicketHandler(logger, ticketStore)

	ticketsHandler.DatabaseName(timeoutContext)
}
