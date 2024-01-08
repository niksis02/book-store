package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/niksis02/book-store/db"
	"github.com/niksis02/book-store/env"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) Register(ctx *fiber.Ctx) error {
	var user db.User
	err := json.Unmarshal(ctx.Body(), &user)
	if err != nil {
		return HandleResponse(ctx, err, nil, &Opts{Status: http.StatusBadRequest, Msg: "Failed to parse the input!"})
	}

	err = c.db.CreateUser(user.Email, user.Password)
	if err != nil {
		return HandleResponse(ctx, err, nil, &Opts{Status: http.StatusBadRequest, Msg: "Failed to create the user account!"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	})

	tkn, err := token.SignedString(env.Env.JWT_SECRET)
	if err != nil {
		return HandleResponse(ctx, err, nil, &Opts{Status: http.StatusBadRequest, Msg: "Failed to create the bearer token!"})
	}

	return HandleResponse(ctx, nil, tkn, &Opts{Status: http.StatusCreated, Msg: "The user has been registered successfully!"})
}

func (c *Controller) Login(ctx *fiber.Ctx) error {
	var input db.User
	err := json.Unmarshal(ctx.Body(), &input)
	if err != nil {
		return HandleResponse(ctx, err, nil, &Opts{Status: http.StatusBadRequest, Msg: "Failed to parse the input!"})
	}

	user, err := c.db.GetUserByEmail(input.Email)
	if err != nil {
		return HandleResponse(ctx, err, nil, &Opts{Status: http.StatusNotFound, Msg: "Incorrect email or password!"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return HandleResponse(ctx, err, nil, &Opts{Status: http.StatusBadRequest, Msg: "Incorrect email or password!"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	})

	tkn, err := token.SignedString(env.Env.JWT_SECRET)
	if err != nil {
		return HandleResponse(ctx, err, nil, &Opts{Status: http.StatusBadRequest, Msg: "Failed to create the bearer token!"})
	}

	return HandleResponse(ctx, nil, tkn, &Opts{Status: http.StatusOK, Msg: "The user has been logged in successfully!"})
}
