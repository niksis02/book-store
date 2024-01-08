package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/niksis02/book-store/controller"
	"github.com/niksis02/book-store/env"
)

var ErrUnauthorized error = errors.New("Unauthorized")

func VerifyJWT(ctx *fiber.Ctx) error {
	// Bearer token authentication indicates that the Authorization header should have
	// the following structure: Bearer <access_token>
	// below comes the validation of the header
	hdr := ctx.Get("Authorization")
	if hdr == "" {
		return controller.HandleResponse(ctx, ErrUnauthorized, nil, &controller.Opts{Status: http.StatusForbidden, Msg: "Authorization header is empty!"})
	}

	tkParts := strings.Split(hdr, " ")
	if len(tkParts) != 2 || tkParts[0] != "Bearer" {
		return controller.HandleResponse(ctx, ErrUnauthorized, nil, &controller.Opts{Status: http.StatusForbidden, Msg: "Invalid Authorization header structure!"})
	}

	tkn := tkParts[1]

	token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return env.Env.JWT_SECRET, nil
	})
	if err != nil {
		return controller.HandleResponse(ctx, err, nil, &controller.Opts{Status: http.StatusForbidden, Msg: "Failed to parse the token!"})
	}

	// Check if token is valid
	if !token.Valid {
		return controller.HandleResponse(ctx, ErrUnauthorized, nil, &controller.Opts{Status: http.StatusForbidden, Msg: "Invalid token"})
	}

	// Access token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return controller.HandleResponse(ctx, ErrUnauthorized, nil, &controller.Opts{Status: http.StatusForbidden, Msg: "Error parsing claims!"})
	}

	// Store user id in the request context in order to query and find the user in controllers
	ctx.Locals("userID", claims["id"].(string))

	return ctx.Next()
}
