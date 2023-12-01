package entity

import (
	"time"
)

type Ticket struct {
<<<<<<< HEAD
	ID          int64  `json:"id"`
	Image       string `json:"image"`
	Location    string `json:"location"`
	Date        string // Format: YYYY-MM-DD
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Status      string `json:"Status"` // e.g., 'available', 'sold'
	Quota       int64  `json:"Quota"`
	Category    string `json:"category"`
	Terjual     int64   `json:"Terjual"` // e.g., 1000, 5000, 10000
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

func NewTicket(image, location, date, title, description, category string, price, quota, terjual int64) *Ticket {
=======
	ID          int64     `json:"id"`
	Image       string    `json:"image"`
	Location    string    `json:"location"`
	Date        string    // Format: YYYY-MM-DD
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Status      string    `json:"Status"` // e.g., 'available', 'sold out'
	Quota       int64     `json:"Quota"`
	Category    string    `json:"category"`
	Tersisa     int64     `json:"Tersisa"` // e.g., 1000, 5000, 10000
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   time.Time `json:"-"`
}

func NewTicket(image, location, date, title, description, category, status string, price, quota, tersisa int64) *Ticket {
>>>>>>> b33534bafa33122c7febb55153d3fad3cbcb800a
	return &Ticket{
		Image:       image,
		Location:    location,
		Date:        date,
		Title:       title,
		Description: description,
		Category:    category,
		Price:       price,
		Status:      status,
		Quota:       quota,
		Terjual:     terjual,
	}
}

func UpdateTicket(id int64, image, location, date, title, description, category string, price, quota, terjual int64) *Ticket {
	return &Ticket{
		ID:          id,
		Image:       image,
		Location:    location,
		Date:        date,
		Title:       title,
		Description: description,
		Category:    category,
		Price:       price,
		Quota:       quota,
		Terjual:     terjual,
	}
}
