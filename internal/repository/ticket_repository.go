package repository
// NOTE :
// FOLDER INI UNTUK MENANGANI KE BAGIAN DATABASE DAN QUERY
import (
	"context"

	"Ticketing/entity"

	"gorm.io/gorm"
)

// ticket repository
type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{
		db: db,
	}
}

// GetAllTickets retrieves all tickets from the database.
func (r *TicketRepository) GetAllTickets(ctx context.Context) ([]*entity.Ticket, error) {
	tickets := make([]*entity.Ticket, 0)
	result := r.db.WithContext(ctx).Find(&tickets)
	if result.Error != nil {
		return nil, result.Error
	}
	return tickets, nil
}

// CreateTicket saves a new ticket to the database.
func (r *TicketRepository) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	result := r.db.WithContext(ctx).Create(&ticket)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateTicket updates a ticket in the database.
func (r *TicketRepository) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	result := r.db.WithContext(ctx).Model(&entity.Ticket{}).Where("id = ?", ticket.ID).Updates(&ticket)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetTicket retrieves a ticket by its ID from the database.
func (r *TicketRepository) GetTicket(ctx context.Context, id int64) (*entity.Ticket, error) {
	ticket := new(entity.Ticket)
	result := r.db.WithContext(ctx).First(&ticket, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return ticket, nil
}

// DeleteTicket deletes a ticket from the database.
func (r *TicketRepository) DeleteTicket(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Delete(&entity.Ticket{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

