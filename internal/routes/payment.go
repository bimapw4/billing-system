package routes

import (
	"billing/internal/handlers"
	"billing/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func PaymentRoutes(app *fiber.App, handler handlers.Handlers, middleware *middleware.Authentication) {
	app.Post("/payment", handler.Payment.PaymentHandler)
	app.Get("/payment", handler.Payment.ListHandler)
}
