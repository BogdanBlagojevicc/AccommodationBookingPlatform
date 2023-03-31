package service

import (
	"ticket-service/repository"
)

type TicketService struct {
	Repo *repository.TicketRepository
}
