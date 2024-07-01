package middleware

import (
	"fmt"

	"github.com/MrAinslay/fiber-rss-feed/internal/auth"
	"github.com/MrAinslay/fiber-rss-feed/packages/config"
	"github.com/MrAinslay/fiber-rss-feed/packages/utils"
	"github.com/gofiber/fiber"
)

type authedHandler func(*fiber.Ctx, config.User)

func MiddlewareAuth(handler authedHandler) fiber.Handler {
	return func(c *fiber.Ctx) {
		apiKey, err := auth.GetApiKey(c)
		if err != nil {
			utils.RespondWithErr(c, 400, fmt.Sprint(err))
			return
		}

		usr, err := config.DBQueris.GetUserById(c.Context(), apiKey)
		if err != nil {
			utils.RespondWithErr(c, 400, fmt.Sprint(err))
			return
		}

		handler(c, usr)
	}
}
