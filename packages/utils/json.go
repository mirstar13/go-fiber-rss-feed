package utils

import "github.com/gofiber/fiber"

func RespondWithJSON(ctx *fiber.Ctx, code int, payload interface{}) {
	rsp := payload

	ctx.Status(code).Send()
	ctx.JSON(rsp)
}

func RespondWithErr(ctx *fiber.Ctx, code int, msg string) {
	type errorRespone struct {
		ErrorMsg string `json:"error"`
	}

	RespondWithJSON(ctx, code, errorRespone{
		ErrorMsg: msg,
	})
}
