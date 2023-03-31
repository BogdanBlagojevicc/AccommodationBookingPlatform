package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type Flight struct {
	ID             primitive.ObjectID
	Departure      string
	DeparturePlace string
	ArrivalPlace   string
	Price          uint64
	NumberOfSeats  uint64
}

func (u *Flight) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}
func (u *Flight) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
