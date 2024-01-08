package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/niksis02/book-store/db"
)

type Controller struct {
	db db.DBService
}

func New(db db.DBService) *Controller {
	return &Controller{
		db: db,
	}
}

type Opts struct {
	Status int
	Msg    string
}

func HandleResponse(ctx *fiber.Ctx, err error, resp any, opts *Opts) error {
	status, msg := 200, ""
	if opts != nil {
		if opts.Status != 0 {
			status = opts.Status
		}
		msg = opts.Msg
	}
	if err != nil {
		return ctx.Status(status).JSON(ErrResponse{Msg: msg, Error: err.Error()})
	}

	return ctx.Status(status).JSON(Response{Data: resp, Msg: msg})
}

type Response struct {
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type ErrResponse struct {
	Msg   string `json:"msg"`
	Error string `json:"error"`
}
