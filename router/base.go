package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/niksis02/book-store/controller"
	"github.com/niksis02/book-store/db"
	"github.com/niksis02/book-store/middleware"
)

type Router struct {
	app *fiber.App
	db  db.DBService
}

func New(app *fiber.App, db db.DBService) *Router {
	return &Router{
		app: app,
		db:  db,
	}
}

// Initialize api endpoints of the application
func (r *Router) Init() {
	ctrl := controller.New(r.db)

	// Initialize request logger
	r.app.Use(logger.New())

	// Public routes
	r.app.Post("/register", ctrl.Register)
	r.app.Post("/login", ctrl.Login)

	// Initialize middlewares
	r.app.Use(middleware.VerifyJWT)

	// Protected routes
	r.app.Put("/order", ctrl.Order)
	r.app.Get("/order", ctrl.ListOrders)
}
