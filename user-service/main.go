package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"user-service/handler"
	"user-service/repository"

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
	userLogger := log.New(os.Stdout, "[user-store] ", log.LstdFlags)

	userStore, err := repository.New(timeoutContext, userLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer userStore.Disconnect(timeoutContext)

	userStore.Ping()

	usersHandler := handler.NewUserHandler(logger, userStore)

	usersHandler.DatabaseName(timeoutContext)

	router := mux.NewRouter()
	router.Use(usersHandler.MiddlewareContentTypeSet)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", usersHandler.PostUser)
	postRouter.Use(usersHandler.MiddlewareUserDeserialization)

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
