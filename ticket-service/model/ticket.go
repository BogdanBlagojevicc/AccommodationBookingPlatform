package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type Ticket struct {
	ID              primitive.ObjectID
	UserID          primitive.ObjectID
	FlightID        primitive.ObjectID
	NumberOfTickets uint8
}

func (u *Ticket) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}
func (u *Ticket) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
