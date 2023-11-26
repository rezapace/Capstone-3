package handler

import (
	"net/http"
	"Ticketing/entity"
	"Ticketing/internal/service"

	"github.com/labstack/echo/v4"
	"strconv"
	"time"
	"Ticketing/internal/http/validator"
)

// TicketHandler handles HTTP requests related to tickets.
type TicketHandler struct {
	ticketService service.TicketUseCase
}

// NewTicketHandler creates a new instance of TicketHandler.
func NewTicketHandler(ticketService service.TicketUseCase) *TicketHandler {
	return &TicketHandler{ticketService}
}

// GetAllTicket 
func (h *TicketHandler) GetAllTickets(c echo.Context) error {
	tickets, err := h.ticketService.GetAllTickets(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}


// CreateTicket
func (h *TicketHandler) CreateTicket(c echo.Context) error {
	var input struct {
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Image       string    `json:"image"`
		Location    string    `json:"location"`
		Date        time.Time `json:"date"`
		Price       float64   `json:"price"`
		Quota       int       `json:"quota"`
	}

	// Input validation
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Convert input.Date to a string with the desired format
	dateStr := input.Date.Format("2006-01-02T15:04:05Z")

	// Create a Ticket object
	ticket := entity.Ticket{
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Location:    input.Location,
		Date:        dateStr, // Assign the formatted date string
		Price:       int64(input.Price),
		Quota:       int64(input.Quota),
		CreatedAt:   time.Now(),
	}

	// Call the ticketService to create the ticket
	err := h.ticketService.CreateTicket(c.Request().Context(), &ticket)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// Return a success message
	return c.JSON(http.StatusCreated, "Ticket created successfully")
}




// GetTicket handles the retrieval of a ticket by ID.
func (h *TicketHandler) GetTicket(c echo.Context) error {
	idStr := c.Param("id") // assuming the ID is passed as a URL parameter as a string
	id, err := strconv.ParseInt(idStr, 10, 64) // Convert the string to int64
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	ticket, err := h.ticketService.GetTicket(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":          ticket.ID,
			"title":       ticket.Title,
			"description": ticket.Description,
			"image":       ticket.Image,
			"location":    ticket.Location,
			"date":        ticket.Date,
			"price":       ticket.Price,
			"quota":       ticket.Quota,
			"created":     ticket.CreatedAt,
		},
	})
}


// UpdateTicket handles the update of an existing ticket.
func (h *TicketHandler) UpdateTicket(c echo.Context) error {
	var input struct {
		ID          int64     `param:"id" validate:"required"`
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Image       string    `json:"image"`
		Location    string    `json:"location"`
		Date        time.Time `json:"date"`
		Price       float64   `json:"price"`
		Quota       int       `json:"quota"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}


	// Convert input.Date to a formatted string
	dateStr := input.Date.Format("2006-01-02T15:04:05Z")

	// Create a Ticket object
	ticket := entity.Ticket{
		ID:          input.ID,            // Assuming ID is already of type int64
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Location:    input.Location,
		Date:        dateStr,            // Assign the formatted date string
		Price:       int64(input.Price),  // Convert Price to int64 if needed
		Quota:       int64(input.Quota),  // Convert Quota to int64 if needed
	}



	err := h.ticketService.UpdateTicket(c.Request().Context(), &ticket)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Ticket updated successfully",
		"ticket":  ticket,
	})
}


// DeleteTicket handles the deletion of a ticket by ID.
func (h *TicketHandler) DeleteTicket(c echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.ticketService.DeleteTicket(c.Request().Context(), input.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Ticket deleted successfully",
	})
}
