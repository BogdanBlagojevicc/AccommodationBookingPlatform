package handler

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"user-service/repository"
)

type UserHandler struct {
	Logger *log.Logger
	Repo   *repository.UserRepository
}

func NewUserHandler(l *log.Logger, r *repository.UserRepository) *UserHandler {
	return &UserHandler{l, r}
}

func (u *UserHandler) DatabaseName(ctx context.Context) {
	dbs, err := u.Repo.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dbs)
}
