package service

import (
	"log"
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

	//flightService.getNumberOfFreeSeats

	newTicket, err := ts.Repo.Insert(ticket)
	if err != nil {
		return nil, err
	}

	//flightService.Update

	return newTicket, nil
}

func (ts *TicketService) GetByUserId(id string) (model.Tickets, error) {
	return ts.Repo.GetByUserId(id)
}
