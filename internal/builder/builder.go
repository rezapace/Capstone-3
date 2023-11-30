package builder

import (
	"Ticketing/internal/config"
	"Ticketing/internal/http/handler"
	"Ticketing/internal/http/router"
	"Ticketing/internal/repository"
	"Ticketing/internal/service"

	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	registrationRepository := repository.NewRegistrationRepository(db)
	registrationService := service.NewRegistrationService(registrationRepository)
	userRepository := repository.NewUserRepository(db) // kenapa make ini? karena nge find email nya dari user_repository
	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)
	authHandler := handler.NewAuthHandler(registrationService, loginService, tokenService)

	return router.PublicRoutes(authHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// Create a user handler
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Create a ticket handler
	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketHandler(ticketService)

	// Create a Blog handler
	BlogRepository := repository.NewBlogRepository(db)
	BlogService := service.NewBlogService(BlogRepository)
	BlogHandler := handler.NewBlogHandler(BlogService)

	// Create an order handler
	OrderRepository := repository.NewOrderRepository(db)
	OrderService := service.NewOrderService(OrderRepository)
	OrderHandler := handler.NewOrderHandler(OrderService)

	NotificationRepository := repository.NewNotificationRepository(db)
	NotificationService := service.NewNotificationService(NotificationRepository)
	NotificationHandler := handler.NewNotificationHandler(NotificationService)

	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(userHandler, ticketHandler, BlogHandler, OrderHandler, NotificationHandler)
}
