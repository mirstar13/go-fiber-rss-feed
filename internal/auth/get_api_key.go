package auth

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

func GetApiKey(ctx *fiber.Ctx) (string, error) {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}

	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}
