package routes

import (
	"billing/internal/handlers"
	"billing/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func BorrowerRoutes(app *fiber.App, handler handlers.Handlers, middleware *middleware.Authentication) {
	app.Post("/borrower", handler.Borrowers.CreateHandler)
	app.Get("/borrower", handler.Borrowers.ListHandler)
}
