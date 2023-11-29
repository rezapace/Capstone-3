package handler

//NOTE :
// FOLDER INI UNTUK MEMANGGIL SERVICE DAN REPOSITORY
import (
	"Ticketing/entity"
	"Ticketing/internal/http/validator"
	"Ticketing/internal/service"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserUsecase
}

// melakukan instace dari user handler
func NewUserHandler(userService service.UserUsecase) *UserHandler {
	return &UserHandler{userService}
}

// func untuk melakukan getAll User
func (h *UserHandler) GetAllUser(ctx echo.Context) error {
	users, err := h.userService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
}

// func untuk melakukan createUser update versi reza v5 halo
func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var input struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email"`
		Number   string `json:"number"`
		Roles    string `json:"roles" validate:"oneof=Admin Buyer"`
		Password string `json:"password"`
	}
	//ini func untuk error checking
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	user := entity.NewUser(input.Name, input.Email, input.Number, input.Roles, input.Password)
	err := h.userService.CreateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	//kalau retrun nya kaya gini akan tampil pesan "User created successfully"
	return ctx.JSON(http.StatusCreated, "User created successfully")
}

// func untuk melakukan updateUser by id
func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var input struct {
		ID       int64  `param:"id" validate:"required"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Number   string `json:"number"`
		Roles    string `json:"roles"`
		Password string `json:"password"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.UpdateUser(input.ID, input.Name, input.Email, input.Number, input.Roles, input.Password)

	err := h.userService.UpdateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "succesfully update user"})
}

// func untuk melakukan getUser by id
func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Jika tidak dapat mengonversi ID menjadi int64, kembalikan respons error
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}
	user, err := h.userService.GetUserByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":       user.ID,
			"name":     user.Name,
			"email":    user.Email,
			"number":   user.Number,
			"password": user.Password,
			"created":  user.CreatedAt,
			"updated":  user.UpdatedAt,
		},
	})
}

// DeleteUser func untuk melakukan delete user by id
func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.userService.Delete(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

// Update User Self
func (h *UserHandler) UpdateUserSelf(ctx echo.Context) error {
	var input struct {
		ID       int64  `param:"id" validate:"required"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Number   string `json:"number"`
		Roles    string `json:"roles" validate:"oneof=Admin Buyer"`
		Password string `json:"password"`
	}

	// Mengambil nilai 'claims' dari JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Mendapatkan nilai 'ID' dari klaim
	userID, ok := claims.Claims.(jwt.MapClaims)["id"].(float64)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user ID from claims")
	}

	// Membandingkan ID yang diterima dari input dengan ID dari klaim
	if int64(userID) != input.ID {
		return ctx.JSON(http.StatusUnprocessableEntity, "you can't update this user")
	}

	// Update user
	user := entity.UpdateUser(input.ID, input.Name, input.Email, input.Number, input.Roles, input.Password)

	// Memanggil service untuk update user
	err := h.userService.UpdateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "successfully update user"})
}
