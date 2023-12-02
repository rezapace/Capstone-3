package handler

//NOTE :
// FOLDER INI UNTUK MEMANGGIL SERVICE DAN REPOSITORY
import (
	"Ticketing/entity"
	"Ticketing/internal/http/validator"
	"Ticketing/internal/service"
	"Ticketing/common"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
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
		Email    string `json:"email" validate:"email"`
		Number   string `json:"number" validate:"min=11,max=13"`
		Roles    string `json:"roles" validate:"oneof=Admin Buyer"`
		Password string `json:"password"`
		Saldo    int64  `json:"saldo"`
	}
	//ini func untuk error checking
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	user := entity.NewUser(input.Name, input.Email, input.Number, input.Roles, input.Password, input.Saldo)
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
		Email    string `json:"email" validate:"email"`
		Number   string `json:"number" validate:"min=11,max=13"`
		Roles    string `json:"roles" validate:"oneof=Admin Buyer"`
		Password string `json:"password"`
		Saldo    int64  `json:"saldo"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.UpdateUser(input.ID, input.Name, input.Email, input.Number, input.Roles, input.Password, input.Saldo)

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
		Email    string `json:"email" validate:"email"`
		Number   string `json:"number" ate:"min=11,max=13"`
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
	user := entity.UpdateUserSelf(input.ID, input.Name, input.Email, input.Number, input.Roles, input.Password)

	// Memanggil service untuk update user
	err := h.userService.UpdateUserSelf(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "successfully update user"})
}

func (h *UserHandler) GetProfile(ctx echo.Context) error {
	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	// Fetch user profile using the user ID
	user, err := h.userService.GetProfile(ctx.Request().Context(), claimsData.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// Return the user profile
	return ctx.JSON(http.StatusOK, user)
}



// delete user self common.JwtCustomClaims
// func (h *UserHandler) DeleteUserSelf(ctx echo.Context) error {
// 	// Pengecekan request
// 	var input struct {
// 		Email string `param:"email" validate:"required,email"`
// 	}

// 	if err := ctx.Bind(&input); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
// 	}

// 	// Mengambil nilai 'claims' dari JWT token
// 	claims, ok := ctx.Get("user").(*jwt.Token)
// 	if !ok {
// 		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
// 	}

// 	// Mendapatkan nilai 'email' dari klaim
// 	jwtClaims, ok := claims.Claims.(*jwt.MapClaims)
// 	if !ok {
// 		return ctx.JSON(http.StatusInternalServerError, "unable to get user email from claims")
// 	}

// 	// Membandingkan email yang diterima dari input dengan email dari klaim
// 	userEmail, ok := (*jwtClaims)["email"].(string)
// 	if !ok {
// 		return ctx.JSON(http.StatusInternalServerError, "unable to get user email from claims")
// 	}

// 	if userEmail != input.Email {
// 		return ctx.JSON(http.StatusUnprocessableEntity, "you can't delete this user")
// 	}

// 	// Delete user
// 	user := entity.DeleteUserSelfByEmail(input.Email)

// 	// Memanggil service untuk delete user
// 	err := h.userService.DeleteUserSelfByEmail(ctx.Request().Context(), user)
// 	if err != nil {
// 		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
// 	}

// 	return ctx.JSON(http.StatusOK, map[string]string{"success": "successfully delete user"})
// }
