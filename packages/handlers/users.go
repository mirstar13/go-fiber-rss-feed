package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/MrAinslay/fiber-rss-feed/packages/config"
	"github.com/MrAinslay/fiber-rss-feed/packages/utils"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HandlerCreateUser(ctx *fiber.Ctx) {
	type parameters struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	type payload struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
		ApiKey    string    `json:"api_key"`
	}
	params := parameters{}

	if err := ctx.BodyParser(&params); err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
	}

	encrPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	usr, err := config.DBQueris.CreateUser(ctx.Context(), config.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Name:      params.Name,
		Password:  string(encrPass),
	})

	log.Println(err)

	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 200, payload{
		Id:        usr.ID,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
		Name:      usr.Name,
		ApiKey:    usr.ApiKey,
	})
}

func HandlerGetUserByKey(ctx *fiber.Ctx, usr config.User) {
	type payload struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
		ApiKey    string    `json:"api_key"`
	}

	usr, err := config.DBQueris.GetUserById(ctx.Context(), usr.ApiKey)
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 200, payload{
		Id:        usr.ID,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
		Name:      usr.Name,
		ApiKey:    usr.ApiKey,
	})
}

func HandlerUserLogin(ctx *fiber.Ctx) {
	type parameters struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	type payload struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
		ApiKey    string    `json:"api_key"`
	}

	params := parameters{}

	if err := ctx.BodyParser(&params); err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	usr, err := config.DBQueris.GetUserByName(ctx.Context(), params.Name)
	if err != nil {
		utils.RespondWithErr(ctx, 401, fmt.Sprint(err))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(params.Password)); err != nil {
		utils.RespondWithErr(ctx, 501, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 201, payload{
		Id:        usr.ID,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
		Name:      usr.Name,
		ApiKey:    usr.ApiKey,
	})
}

func HandlerUpdateUser(ctx *fiber.Ctx, usr config.User) {
	type parameters struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	type payload struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
		ApiKey    string    `json:"api_key"`
	}

	params := parameters{}
	if err := ctx.BodyParser(&params); err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	if params.Name == "" {
		params.Name = usr.Name
	}

	var encrPass []byte
	var err error
	if params.Password == "" {
		params.Password = usr.Password
		encrPass = []byte(params.Password)
	} else {
		encrPass, err = bcrypt.GenerateFromPassword([]byte(params.Password), 10)
		if err != nil {
			utils.RespondWithErr(ctx, 401, fmt.Sprint(err))
			return
		}
	}

	upUsr, err := config.DBQueris.UpdateUser(ctx.Context(), config.UpdateUserParams{
		Name:      params.Name,
		Password:  string(encrPass),
		UpdatedAt: time.Now(),
		ApiKey:    usr.ApiKey,
	})
	if err != nil {
		utils.RespondWithErr(ctx, 401, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 201, payload{
		Id:        upUsr.ID,
		CreatedAt: upUsr.CreatedAt,
		UpdatedAt: upUsr.UpdatedAt,
		Name:      upUsr.Name,
		ApiKey:    upUsr.ApiKey,
	})
}

func HandlerDeleteUser(ctx *fiber.Ctx, usr config.User) {
	type deleteResponse struct {
		Message string `json:"message"`
	}

	config.DBQueris.DeleteUser(ctx.Context(), usr.ApiKey)

	utils.RespondWithJSON(ctx, 201, deleteResponse{
		Message: fmt.Sprintf("Successfully deleted user with id %s", usr.ID),
	})
}
