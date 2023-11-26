package router

import (
	"Ticketing/internal/http/handler"

	"github.com/labstack/echo/v4"
)

// membuat struct route
type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
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
func PrivateRoutes(UserHandler *handler.UserHandler, TicketHandler *handler.TicketHandler, BlogHandler *handler.BlogHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: UserHandler.CreateUser,
		},

		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: UserHandler.GetAllUser,
		},

		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: UserHandler.UpdateUser,
		},

		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: UserHandler.GetUserByID,
		},

		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: UserHandler.DeleteUser,
		},
		{
			Method:  echo.POST,
			Path:    "/ticket",
			Handler: TicketHandler.CreateTicket,
		},

		{
			Method:  echo.GET,
			Path:    "/ticket",
			Handler: TicketHandler.GetAllTickets,
		},

		{
			Method:  echo.PUT,
			Path:    "/ticket/:id",
			Handler: TicketHandler.UpdateTicket,
		},

		{
			Method:  echo.GET,
			Path:    "/ticket/:id",
			Handler: TicketHandler.GetTicket,
		},

		{
			Method:  echo.DELETE,
			Path:    "/ticket/:id",
			Handler: TicketHandler.DeleteTicket,
		},

		{
			Method:  echo.GET,
			Path:    "/ticket/search/:search",
			Handler: TicketHandler.SearchTicket,
		},

		{
			Method:  echo.POST,
			Path:    "/blog",
			Handler: BlogHandler.CreateBlog,
		},

		{
			Method:  echo.GET,
			Path:    "/blog",
			Handler: BlogHandler.GetAllBlogs,
		},

		{
			Method:  echo.PUT,
			Path:    "/blog/:id",
			Handler: BlogHandler.UpdateBlog,
		},

		{
			Method:  echo.GET,
			Path:    "/blog/:id",
			Handler: BlogHandler.GetBlog,
		},

		{
			Method:  echo.DELETE,
			Path:    "/blog/:id",
			Handler: BlogHandler.DeleteBlog,
		},

		{
			Method:  echo.GET,
			Path:    "/blog/search/:search",
			Handler: BlogHandler.SearchBlog,
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
