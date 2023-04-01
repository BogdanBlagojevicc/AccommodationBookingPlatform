package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type Ticket struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	UserID   primitive.ObjectID `bson:"userId" json:"userId"`
	FlightID primitive.ObjectID `bson:"flightId" json:"flightId"`
}

func (u *Ticket) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}
func (u *Ticket) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
