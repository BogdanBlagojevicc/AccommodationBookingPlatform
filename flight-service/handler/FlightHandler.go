package handler

import (
	"context"
	"flight-service/model"
	"flight-service/service"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strconv"
)

type KeyProduct struct{}
type FlightHandler struct {
	Logger  *log.Logger
	Service *service.FlightService
}

func NewFlightHandler(l *log.Logger, s *service.FlightService) *FlightHandler {
	return &FlightHandler{l, s}
}

func (u *FlightHandler) DatabaseName(ctx context.Context) {
	dbs, err := u.Service.Repo.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dbs)
}

func (u *FlightHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		flight := &model.Flight{}
		err := flight.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.Logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, flight)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (f *FlightHandler) PostFlight(rw http.ResponseWriter, h *http.Request) {
	flight := h.Context().Value(KeyProduct{}).(*model.Flight)
	//newUser := model.User{ID: primitive.NewObjectID(), FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Password: user.Password}
	flight.ID = primitive.NewObjectID()
	createdFlight, err := f.Service.Insert(flight)
	if createdFlight == nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusCreated)
}

func (f *FlightHandler) GetFlights(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	date := vars["departure"]
	departurePlace := vars["departurePlace"]
	arrivalPlace := vars["arrivalPlace"]
	noOfSeats := vars["noOfSeats"]

	n, erre := strconv.Atoi(noOfSeats)
	if erre != nil {
		fmt.Println("Error during conversion.")
		return
	}
	fmt.Println("Broj iz URLA: ", n)

	flights, err := f.Service.GetFlights(date, departurePlace, arrivalPlace, n)

	if err != nil {
		f.Logger.Println("Database exception ", err)
	}

	if flights == nil {
		return
	}
	err = flights.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		f.Logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (u *FlightHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		u.Logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
