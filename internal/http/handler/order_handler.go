package handler

import (
	"Ticketing/entity"
	"Ticketing/internal/http/validator"
	"Ticketing/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService service.OrderUsecase
}

func NewOrderHandler(Orderservice *service.OrderService) *OrderHandler {
	return &OrderHandler{Orderservice}
}

// func untuk create order
func (h *OrderHandler) CreateOrder(ctx echo.Context) error {
	var input struct {
		TicketID int64 `json:"ticket_id" validate:"required"`
		Quantity int64 `json:"quantity" validate:"required"`
		UserID   int64 `json:"user_id" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	order := entity.NewOrder(input.TicketID, input.Quantity, input.UserID)
	err := h.orderService.CreateOrder(ctx.Request().Context(), order)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"OrderAt": order.OrderAt,
	})
}

// Get All Order
func (h *OrderHandler) GetAllOrders(ctx echo.Context) error {
	// Implementasi untuk mendapatkan semua pesanan dari usecase
	orders, err := h.orderService.GetOrders(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	var orderDetails []map[string]interface{}
	for _, order := range orders {
		ticket, err := h.orderService.GetTicketByID(ctx.Request().Context(), order.TicketID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusInternalServerError, err.Error()))
		}

		orderDetail := map[string]interface{}{
			"user_id": order.UserID,
			"ticket":  ticket,
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Get all orders success",
		"order_details": orderDetails,
	})
}

// get order by user_id
func (h *OrderHandler) GetOrderByUserID(ctx echo.Context) error {
	// Implementasi untuk mendapatkan semua pesanan dari usecase
	orders, err := h.orderService.GetOrders(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	var orderDetails []map[string]interface{}
	for _, order := range orders {
		ticket, err := h.orderService.GetTicketByID(ctx.Request().Context(), order.TicketID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusInternalServerError, err.Error()))
		}

		orderDetail := map[string]interface{}{
			"user_id": order.UserID,
			"ticket":  ticket,
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Get all orders success",
		"order_details": orderDetails,
	})
}
