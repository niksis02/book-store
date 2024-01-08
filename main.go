package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/niksis02/book-store/db"
	"github.com/niksis02/book-store/env"
	"github.com/niksis02/book-store/router"
)

func main() {
	app := fiber.New()

	err := env.LoadEnvVars()
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := db.NewPostgres()
	if err != nil {
		log.Fatalln(err)
	}

	r := router.New(app, conn)
	r.Init()
	if err := app.Listen(fmt.Sprintf(":%v", env.Env.PORT)); err != nil {
		log.Fatalln(err)
	}
}
