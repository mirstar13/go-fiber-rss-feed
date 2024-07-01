package handlers

import (
	"fmt"
	"time"

	"github.com/MrAinslay/fiber-rss-feed/packages/config"
	"github.com/MrAinslay/fiber-rss-feed/packages/models"
	"github.com/MrAinslay/fiber-rss-feed/packages/utils"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func HandlerCreateFeed(ctx *fiber.Ctx, usr config.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	type payload struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		UserId    uuid.UUID `json:"user_id"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
	}

	params := parameters{}
	if err := ctx.BodyParser(&params); err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	feed, err := config.DBQueris.CreateFeed(ctx.Context(), config.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UserID:    usr.ID,
		Name:      params.Name,
		Url:       params.URL,
	})
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 200, payload{
		Id:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		UserId:    usr.ID,
		Name:      feed.Name,
		URL:       feed.Url,
	})
}

func HandlerGetFeeds(ctx *fiber.Ctx) {
	feeds, err := config.DBQueris.GetAllFeeds(ctx.Context())
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 200, models.DatabaseFeedsToFeeds(feeds))
}

func HandlerGetFeedById(ctx *fiber.Ctx) {
	uuid, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		utils.RespondWithErr(ctx, 401, fmt.Sprint(err))
		return
	}

	feed, err := config.DBQueris.GetFeedById(ctx.Context(), uuid)
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 201, models.DatabaseFeedToFeed(feed))
}

func HandlerDeleteFeed(ctx *fiber.Ctx, usr config.User) {
	type deleteResponse struct {
		Message string `json:"message"`
	}

	feedUUID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	feed, err := config.DBQueris.GetFeedById(ctx.Context(), feedUUID)
	if err != nil {
		utils.RespondWithErr(ctx, 401, fmt.Sprint(err))
		return
	}

	if feed.UserID != usr.ID {
		utils.RespondWithErr(ctx, 501, "Not authorized")
		return
	}

	config.DBQueris.DeleteFeed(ctx.Context(), feed.ID)
	utils.RespondWithJSON(ctx, 201, deleteResponse{
		Message: fmt.Sprintf("Successfully deleted feed with id %s", feedUUID),
	})
}
