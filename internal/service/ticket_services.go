package service

import (
	"context"

	"Ticketing/entity"
)

// TicketUseCase is an interface for ticket-related use cases.
type TicketUseCase interface {
	GetAllTickets(ctx context.Context) ([]*entity.Ticket, error)
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	GetTicket(ctx context.Context, id int64) (*entity.Ticket, error)
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	SearchTicket(ctx context.Context, search string) ([]*entity.Ticket, error)
	DeleteTicket(ctx context.Context, id int64) error
}

type TicketRepository interface {
	GetAllTickets(ctx context.Context) ([]*entity.Ticket, error)
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	GetTicket(ctx context.Context, id int64) (*entity.Ticket, error)
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	SearchTicket(ctx context.Context, search string) ([]*entity.Ticket, error)
	DeleteTicket(ctx context.Context, id int64) error
}

// TicketService is responsible for ticket-related business logic.
type TicketService struct {
	Repository TicketRepository
}

// NewTicketService creates a new instance of TicketService.
func NewTicketService(Repository TicketRepository) *TicketService {
	return &TicketService{Repository: Repository}
}

func (s *TicketService) GetAllTickets(ctx context.Context) ([]*entity.Ticket, error) {
	return s.Repository.GetAllTickets(ctx)
}

func (s *TicketService) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.Repository.CreateTicket(ctx, ticket)
}

func (s *TicketService) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.Repository.UpdateTicket(ctx, ticket)
}

func (s *TicketService) GetTicket(ctx context.Context, id int64) (*entity.Ticket, error) {
	return s.Repository.GetTicket(ctx, id)
}

func (s *TicketService) DeleteTicket(ctx context.Context, id int64) error {
	return s.Repository.DeleteTicket(ctx, id)
}

// search ticket
func (s *TicketService) SearchTicket(ctx context.Context, search string) ([]*entity.Ticket, error) {
	return s.Repository.SearchTicket(ctx, search)
}