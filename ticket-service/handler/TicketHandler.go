package handler

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"ticket-service/model"
	"ticket-service/repository"
)

type KeyProduct struct{}

type TicketHandler struct {
	Logger *log.Logger
	Repo   *repository.TicketRepository
}

func NewTicketHandler(l *log.Logger, r *repository.TicketRepository) *TicketHandler {
	return &TicketHandler{l, r}
}

func (u *TicketHandler) DatabaseName(ctx context.Context) {
	dbs, err := u.Repo.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dbs)
}

func (u *TicketHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		ticket := &model.Ticket{}
		err := ticket.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.Logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, ticket)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (u *TicketHandler) PostTicket(rw http.ResponseWriter, h *http.Request) {

}

func (u *TicketHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		u.Logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
