package model

import "github.com/google/uuid"

type Ticket struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	FlightID        uuid.UUID
	NumberOfTickets uint8
}
