package handler

import (
    "Ticketing/entity"
    "Ticketing/internal/service"
    "net/http"
    "github.com/labstack/echo/v4"
)

type TopupHandler struct {
    topupService service.TopupService
}

func NewTopupHandler(topupService service.TopupService) *TopupHandler {
    return &TopupHandler{topupService}
}

func (h *TopupHandler) CreateTopup(c echo.Context) error {
    var topup entity.Topup
    if err := c.Bind(&topup); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    chargeResp, err := h.topupService.CreateMidtransCharge(topup.ID, int64(topup.Amount))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    topup.SnapURL = chargeResp.RedirectURL

    // Perhatikan penambahan c.Request().Context() di sini
    newTopup, err := h.topupService.CreateTopup(c.Request().Context(), topup)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, newTopup)
}


// topup saldo sederhana
// func (h *TopupHandler) TopupSaldo(c echo.Context) error {
//     var topup entity.Topup
//     if err := c.Bind(&topup); err != nil {
//         return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
//     }

//     newTopup, err := h.topupService.TopupSaldo(c.Request().Context(), topup)
//     if err != nil {
//         return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
//     }

//     return c.JSON(http.StatusCreated, newTopup)
// }