package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type TicketRepository struct {
	Cli    *mongo.Client
	Logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*TicketRepository, error) {

	dburi := os.Getenv("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &TicketRepository{
		Cli:    client,
		Logger: logger,
	}, nil
}
func (u *TicketRepository) Disconnect(ctx context.Context) error {
	err := u.Cli.Disconnect(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *TicketRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.Cli.Ping(ctx, readpref.Primary())
	if err != nil {
		u.Logger.Println(err)
	}

	dbs, err := u.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		u.Logger.Println(err)
	}
	fmt.Println(dbs)
}