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

func HandlerCreatePostLike(ctx *fiber.Ctx, usr config.User) {
	type parameters struct {
		PostId uuid.UUID `json:"post_id"`
	}

	params := parameters{}
	if err := ctx.BodyParser(&params); err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	like, err := config.DBQueris.CreatePostLike(ctx.Context(), config.CreatePostLikeParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UserID:    usr.ID,
		PostID:    params.PostId,
	})
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 200, models.DatabasePostLikeToPostLike(like))
}

func HandlerGetPostLikesByUser(ctx *fiber.Ctx, usr config.User) {
	likes, err := config.DBQueris.GetPostLikesByUser(ctx.Context(), usr.ID)
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 201, models.DatabasePostLikesToPostLikes(likes))
}

func HandlerDeletePostLike(ctx *fiber.Ctx, usr config.User) {
	type deleteResponse struct {
		Message string `json:"message"`
	}

	likeUUID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	like, err := config.DBQueris.GetPostLikeById(ctx.Context(), likeUUID)
	if err != nil {
		utils.RespondWithErr(ctx, 401, fmt.Sprint(err))
		return
	}

	if like.UserID != usr.ID {
		utils.RespondWithErr(ctx, 501, "Not Authorized")
		return
	}

	config.DBQueris.DeletePostLike(ctx.Context(), likeUUID)
	utils.RespondWithJSON(ctx, 201, deleteResponse{
		Message: fmt.Sprintf("successfully deleted post like with id %s", likeUUID),
	})
}
