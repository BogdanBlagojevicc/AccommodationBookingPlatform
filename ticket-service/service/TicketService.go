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

func NewTicketService(l *log.Logger, s *repository.TicketRepository) *TicketService {
	return &TicketService{l, s}
}

func (ts *TicketService) Insert(newTicket *model.Ticket) (*model.Ticket, error) {
	flightId := newTicket.FlightID
	ts.Logger.Printf("TICKEEEEET: " + flightId.String())
	//n := fs.FlightSer.GetNumberOfFreeSeatsById(flightId)
	//n := service.FlightService{}

	// ovako nesto
	//req_url := fmt.Sprintf("http://%s:%s/getNumberOfFreeSeatsById/%s", os.Getenv("FLIGHT_SERVICE_DOMAIN"), os.Getenv("FLIGHT_SERVICE_PORT"), flightId)
	//json_orders, _ := json.Marshal(newTicket)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", req_url)
	//fmt.Println(string(json_orders))
	////resp, err := http.Post(req_url, "application/json", bytes.NewBuffer(json_orders))
	//resp, err := http.Get(req_url)
	//if err != nil || resp.StatusCode == 404 {
	//	print("Failed creating ticket in flight-service")
	//}

	return ts.Repo.Insert(newTicket)
}
