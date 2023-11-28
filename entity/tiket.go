package entity

import (
	"time"
)

type Ticket struct {
	ID          int64
	Image       string
	Location    string
	Date        string // Format: YYYY-MM-DD
	Title       string
	Description string
	Price       int64
	Status      string // e.g., 'available', 'sold out'
	Quota       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewTicket(image, location, date, title, description string, price, quota int64) *Ticket {
	return &Ticket{
		Image:       image,
		Location:    location,
		Date:        date,
		Title:       title,
		Description: description,
		Price:       price,
		Quota:       quota,
		CreatedAt:   time.Now(),
	}
}

func UpdateTicket(id int64, image, location, date, title, description string, price, quota int64) *Ticket {
	return &Ticket{
		ID:          id,
		Image:       image,
		Location:    location,
		Date:        date,
		Title:       title,
		Description: description,
		Price:       price,
		Quota:       quota,
		UpdatedAt:   time.Now(),
	}
}
