package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"ticket-service/handler"
	"ticket-service/repository"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	ticketLogger := log.New(os.Stdout, "[ticket-store] ", log.LstdFlags)

	ticketStore, err := repository.New(timeoutContext, ticketLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer ticketStore.Disconnect(timeoutContext)

	ticketStore.Ping()

	ticketsHandler := handler.NewTicketHandler(logger, ticketStore)

	ticketsHandler.DatabaseName(timeoutContext)

	router := mux.NewRouter()
	router.Use(ticketsHandler.MiddlewareContentTypeSet)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	//postRouter.HandleFunc("/", ticketsHandler.PostFlight)
	postRouter.Use(ticketsHandler.MiddlewareUserDeserialization)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")

}
