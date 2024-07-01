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

func HandlerGetUserFeedFollows(ctx *fiber.Ctx, usr config.User) {
	follows, err := config.DBQueris.GetFeedFollows(ctx.Context(), usr.ID)
	if err != nil {
		utils.RespondWithErr(ctx, 401, fmt.Sprint(err))
	}

	utils.RespondWithJSON(ctx, 201, models.DatabaseFeedFollowsToFeedFollows(follows))
}

func HandlerCreateFeedFollow(ctx *fiber.Ctx, usr config.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	params := parameters{}
	if err := ctx.BodyParser(&params); err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	feedFollow, err := config.DBQueris.CreateFeedFollow(ctx.Context(), config.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UserID:    usr.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 200, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func HandlerDeleteFeedFollow(ctx *fiber.Ctx, usr config.User) {
	type deleteResponse struct {
		Message string `json:"message"`
	}

	followUUID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	follow, err := config.DBQueris.GetFeedById(ctx.Context(), followUUID)
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	if follow.UserID != usr.ID {
		utils.RespondWithErr(ctx, 501, "Not authorized")
		return
	}

	config.DBQueris.DeleteFeed(ctx.Context(), followUUID)
	utils.RespondWithJSON(ctx, 201, deleteResponse{
		Message: fmt.Sprintf("Successfully deleted feed follow with id %s", followUUID),
	})
}
