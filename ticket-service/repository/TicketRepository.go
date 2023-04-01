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
	"ticket-service/model"
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

func (tr *TicketRepository) getCollection() *mongo.Collection {
	bookingDatabase := tr.Cli.Database("booking")
	ticketsCollection := bookingDatabase.Collection("tickets")
	return ticketsCollection
}

func (tr *TicketRepository) Insert(ticket *model.Ticket) (*model.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ticketsCollection := tr.getCollection()

	result, err := ticketsCollection.InsertOne(ctx, &ticket)
	if err != nil {
		tr.Logger.Println(err)
		return nil, err
	}
	tr.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	return ticket, nil
}
