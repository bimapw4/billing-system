package routes

import (
	"billing/internal/handlers"
	"billing/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func LoanRoutes(app *fiber.App, handler handlers.Handlers, middleware *middleware.Authentication) {
	app.Post("/loan", handler.Loan.CreateHandler)
	app.Get("/loan", handler.Loan.ListHandler)
}
