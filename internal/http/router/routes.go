package router

import (
	"Ticketing/internal/http/handler"

	"github.com/labstack/echo/v4"
)

const (
	Admin = "Admin"
	Buyer = "Buyer"
)

var (
	allRoles  = []string{Admin, Buyer}
	onlyAdmin = []string{Admin}
	onlyBuyer = []string{Buyer}
)

// membuat struct route
type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
	Role    []string
}

// membuat fungsi untuk mengembalikan route
// pada func ini perlu login krna private
func PublicRoutes(authHandler *handler.AuthHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: authHandler.Login,
		},
		{
			Method:  echo.POST,
			Path:    "/register",
			Handler: authHandler.Registration,
		},
	}
}

// membuat fungsi untuk mengembalikan route
// pada func ini tdk perlu login krna public
func PrivateRoutes(UserHandler *handler.UserHandler, TicketHandler *handler.TicketHandler, BlogHandler *handler.BlogHandler, OrderHandler *handler.OrderHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: UserHandler.CreateUser,
			Role:    allRoles,
		},

		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: UserHandler.GetAllUser,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: UserHandler.UpdateUser,
			Role:    allRoles,
		},

		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: UserHandler.GetUserByID,
			Role:    allRoles,
		},

		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: UserHandler.DeleteUser,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.PUT,
			Path:    "/users/self",
			Handler: UserHandler.UpdateUserSelf,
			Role:    allRoles,
		},
		{
			Method:  echo.POST,
			Path:    "/ticket",
			Handler: TicketHandler.CreateTicket,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.GET,
			Path:    "/ticket",
			Handler: TicketHandler.GetAllTickets,
			Role:    allRoles,
		},

		{
			Method:  echo.PUT,
			Path:    "/ticket/:id",
			Handler: TicketHandler.UpdateTicket,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.GET,
			Path:    "/ticket/:id",
			Handler: TicketHandler.GetTicket,
			Role:    allRoles,
		},

		{
			Method:  echo.DELETE,
			Path:    "/ticket/:id",
			Handler: TicketHandler.DeleteTicket,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.GET,
			Path:    "/ticket/search/:search",
			Handler: TicketHandler.SearchTicket,
			Role:    allRoles,
		},

		{
			Method:  echo.POST,
			Path:    "/blog",
			Handler: BlogHandler.CreateBlog,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.GET,
			Path:    "/blog",
			Handler: BlogHandler.GetAllBlogs,
			Role:    allRoles,
		},

		{
			Method:  echo.PUT,
			Path:    "/blog/:id",
			Handler: BlogHandler.UpdateBlog,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.GET,
			Path:    "/blog/:id",
			Handler: BlogHandler.GetBlog,
			Role:    allRoles,
		},

		{
			Method:  echo.DELETE,
			Path:    "/blog/:id",
			Handler: BlogHandler.DeleteBlog,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.GET,
			Path:    "/blog/search/:search",
			Handler: BlogHandler.SearchBlog,
			Role:    allRoles,
		},

		{
			Method:  echo.POST,
			Path:    "/order",
			Handler: OrderHandler.CreateOrder,
			Role:    onlyBuyer,
		},

		{
			Method:  echo.GET,
			Path:    "/order",
			Handler: OrderHandler.GetAllOrders,
			Role:    allRoles,
		},

		{
			Method:  echo.GET,
			Path:    "/order/:id",
			Handler: OrderHandler.GetOrderByUserID,
			Role:    allRoles,
		},

		//filter ticket by location
		{
			Method:  echo.GET,
			Path:    "/ticket/location/:location",
			Handler: TicketHandler.FilterTicket,
			Role:    allRoles,
		},
		// filter ticket by category
		{
			Method:  echo.GET,
			Path:    "/ticket/category/:category",
			Handler: TicketHandler.FilterTicketByCategory,
			Role:    allRoles,
		},
		// filter ticket by range time (start - end)
		{
			Method:  echo.GET,
			Path:    "/ticket/range/:start/:end",
			Handler: TicketHandler.FilterTicketByRangeTime,
			Role:    allRoles,
		},
		// filter ticket by price (min - max)
		{
			Method:  echo.GET,
			Path:    "/ticket/price/:min/:max",
			Handler: TicketHandler.FilterTicketByPrice,
			Role:    allRoles,
		},
	}
}

//NOTE :
//MENGAPA TERDAPAT 2 FUNC DIATAS? YAITU PUBLIC DAN PRIVATE
//KAREN DI SERVER.GO KITA BUAT GROUP API, DAN KITA MEMBAGI ROUTE YANG PERLU LOGIN DAN TIDAK PERLU LOGIN
// YAITU PUBLIC DAN PRIVATE

//note ;
//untuk menjalankan nya setelah port 8080 ditambahin /api/v1
// karna di server.go kita membuat group API
