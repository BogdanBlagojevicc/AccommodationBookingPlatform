package handler

import (
	"context"
	"flight-service/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

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
