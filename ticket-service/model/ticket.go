package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	ID              primitive.ObjectID
	UserID          primitive.ObjectID
	FlightID        primitive.ObjectID
	NumberOfTickets uint8
}
