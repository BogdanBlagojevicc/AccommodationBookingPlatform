package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	ID             primitive.ObjectID
	Departure      string
	DeparturePlace string
	ArrivalPlace   string
	Price          uint64
	NumberOfSeats  uint64
}
