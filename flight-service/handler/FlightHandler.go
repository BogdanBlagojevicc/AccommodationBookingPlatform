package handler

import (
	"context"
	"flight-service/model"
	"flight-service/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type KeyProduct struct{}
type FlightHandler struct {
	Logger *log.Logger
	Repo   *repository.FlightRepository
}

func NewFlightHandler(l *log.Logger, r *repository.FlightRepository) *FlightHandler {
	return &FlightHandler{l, r}
}

func (u *FlightHandler) DatabaseName(ctx context.Context) {
	dbs, err := u.Repo.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dbs)
}

func (u *FlightHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		flight := &model.Flight{}
		err := flight.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.Logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, flight)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (u *FlightHandler) PostFlight(rw http.ResponseWriter, h *http.Request) {
	flight := h.Context().Value(KeyProduct{}).(*model.Flight)
	flight.ID = primitive.NewObjectID()
	u.Repo.Insert(flight)
	rw.WriteHeader(http.StatusCreated)
}

func (u *FlightHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		u.Logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
