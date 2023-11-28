package repository

import (
	"Ticketing/entity"
	"context"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *entity.Order) error {
	err := r.db.WithContext(ctx).Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTicket(ctx context.Context, ticketID int64) (*entity.Ticket, error) {
	ticket := new(entity.Ticket)
	if err := r.db.WithContext(ctx).Where("id = ?", ticketID).First(&ticket).Error; err != nil {
		return nil, err
	}
	return ticket, nil
}

func (r *OrderRepository) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.Ticket{}).
		Where("id = ?", ticket.ID).
		Updates(&ticket).Error; err != nil {
		return err
	}
	return nil
}

// Add the following method to implement the missing GetTicketByID
func (r *OrderRepository) GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error) {
	ticket := new(entity.Ticket)
	result := r.db.WithContext(ctx).First(&ticket, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return ticket, nil
}

// repository order.go
func (r *OrderRepository) GetOrders(ctx context.Context) ([]*entity.Order, error) {
	orders := make([]*entity.Order, 0)
	err := r.db.WithContext(ctx).Preload("Ticket").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// get order by user_id
func (r *OrderRepository) GetOrderByUserID(ctx context.Context, userID int64) ([]*entity.Order, error) {
	orders := make([]*entity.Order, 0)
	err := r.db.WithContext(ctx).Preload("Ticket").Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}