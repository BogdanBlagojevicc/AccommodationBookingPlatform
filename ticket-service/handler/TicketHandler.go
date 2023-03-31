package handler

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"ticket-service/repository"
)

type TicketHandler struct {
	Logger *log.Logger
	Repo   *repository.TicketRepository
}

func NewTicketHandler(l *log.Logger, r *repository.TicketRepository) *TicketHandler {
	return &TicketHandler{l, r}
}

func (u *TicketHandler) DatabaseName(ctx context.Context) {
	dbs, err := u.Repo.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dbs)
}
