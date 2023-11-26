package server

import (
	"Ticketing/internal/config"
	"Ticketing/internal/http/binder"
	"Ticketing/internal/http/router"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// merupakan struct dari eco
type Server struct {
	*echo.Echo
}

// untuk membuat server
func NewServer(
	cfg *config.Config,
	binder *binder.Binder,
	publicRoutes, privateRoutes []*router.Route) *Server {
	e := echo.New()
	e.HideBanner = true // untuk menghilangkan banner echo, karena sudah membuat banner sendiri di splash
	e.Binder = binder

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
	)

	//membuat group API
	v1 := e.Group("/api/v1")

	for _, public := range publicRoutes {
		//e.add = untuk menambahkan route baru.
		v1.Add(public.Method, public.Path, public.Handler)
	}

	//ketika sudah ingin menggunakan middleware, maka menambahkan private.Middleware.
	for _, private := range privateRoutes {
		v1.Add(private.Method, private.Path, private.Handler, JWTprotected(cfg.JWT.SecretKey))
	}

	//hedler untuk mengecek kesehatan server
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	return &Server{e}

}

// func untuk pendeklarasian JWT Middleware
func JWTprotected(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secretKey),
	})
}
