package routes

import (
	"billing/internal/handlers"
	"billing/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication) {
	// register route
	routes := []func(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication){
		BorrowerRoutes,
		LoanRoutes,
	}

	for _, route := range routes {
		route(app, handler, m)
	}
}
