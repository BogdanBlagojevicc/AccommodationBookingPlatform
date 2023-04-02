package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"ticket-service/model"
	"ticket-service/repository"
)

type TicketService struct {
	Logger *log.Logger
	Repo   *repository.TicketRepository
}

func NewTicketService(l *log.Logger, r *repository.TicketRepository) *TicketService {
	return &TicketService{Logger: l, Repo: r}
}
func (ts *TicketService) Insert(ticket *model.Ticket) (*model.Ticket, error) {

	reqUrl := fmt.Sprintf("http://%s:%s/%s/%s", os.Getenv("FLIGHT_SERVICE_DOMAIN"), os.Getenv("FLIGHT_SERVICE_PORT"), ticket.FlightID, ticket.NumberOfTickets)
	fmt.Printf("Sending GET request to url %s\n", reqUrl)

	resp, err := http.Get(reqUrl)
	if err != nil || resp.StatusCode == 400 {
		ts.Logger.Println("Failed")
		return nil, err
	}

	newTicket, err := ts.Repo.Insert(ticket)
	if err != nil {
		return nil, err
	}

	return newTicket, nil
}

func (ts *TicketService) GetByUserId(id string) (model.Tickets, error) {
	return ts.Repo.GetByUserId(id)
}
