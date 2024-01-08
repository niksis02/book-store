package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Books struct {
	BookIds []string `json:"bookIds"`
}

func (c *Controller) Order(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(string)
	var books Books
	err := json.Unmarshal(ctx.Body(), &books)
	if err != nil {
		return HandleResponse(ctx, err, nil, &Opts{Status: http.StatusBadRequest, Msg: "Failed to parse the input!"})
	}

	err = c.db.CreateOrders(userID, books.BookIds)
	if err != nil {
		return HandleResponse(ctx, err, nil, &Opts{Status: http.StatusBadRequest, Msg: "Failed to place the orders!"})
	}

	return HandleResponse(ctx, nil, "", &Opts{Status: http.StatusCreated, Msg: "The orders have successfully been created!"})
}

func (c Controller) ListOrders(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(string)

	orders, err := c.db.ListUserOrders(userID)
	if err != nil {
		return HandleResponse(ctx, err, nil, &Opts{Status: http.StatusBadRequest, Msg: "Failed to list the orders!"})
	}

	return HandleResponse(ctx, nil, orders, nil)
}
