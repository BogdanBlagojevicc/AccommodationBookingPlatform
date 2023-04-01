package repository

import (
	"context"
	"flight-service/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type FlightRepository struct {
	Cli    *mongo.Client
	Logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*FlightRepository, error) {

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
	return &FlightRepository{
		Cli:    client,
		Logger: logger,
	}, nil
}
func (u *FlightRepository) Disconnect(ctx context.Context) error {
	err := u.Cli.Disconnect(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *FlightRepository) Ping() {
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

func (ur *FlightRepository) getCollection() *mongo.Collection {
	bookingDatabase := ur.Cli.Database("booking")
	usersCollection := bookingDatabase.Collection("flights")
	return usersCollection
}

func (ur *FlightRepository) Insert(flight *model.Flight) (*model.Flight, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flightsCollection := ur.getCollection()

	result, err := flightsCollection.InsertOne(ctx, &flight)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}
	ur.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	return flight, nil
}
