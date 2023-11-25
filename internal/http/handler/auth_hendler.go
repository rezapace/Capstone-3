package handler

import (
	"Ticketing/internal/http/validator"
	"Ticketing/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	loginService service.LoginUseCase //untuk memanggil service yang ngelakuin pengecekan user.
	tokenService service.TokenUsecase //untuk memanggil func akses token
}

// ini func untuk type AuthHandler
func NewAuthHandler(
	loginService service.LoginUseCase,
	tokenService service.TokenUsecase,
) *AuthHandler {
	return &AuthHandler{
		loginService: loginService,
		tokenService: tokenService,
	}
}

// func ini untuk login
func (h *AuthHandler) Login(ctx echo.Context) error {
	//pengecekan request
	var input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	if err := ctx.Bind(&input); err != nil { // di cek pake validate buat masukin input
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	//untuk manggil login service di folder service
	user, err := h.loginService.Login(ctx.Request().Context(), input.Email, input.Password)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	//untuk manggil token service di folder service
	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	data := map[string]interface{}{
		"access_token": accessToken,
	}
	return ctx.JSON(http.StatusOK, data)
}
