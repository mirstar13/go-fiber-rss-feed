package routes

import (
	"github.com/MrAinslay/fiber-rss-feed/packages/handlers"
	"github.com/MrAinslay/fiber-rss-feed/packages/middleware"
	"github.com/gofiber/fiber"
)

var RegisterFeedRoutes = func(app *fiber.App) {
	app.Get("/v1/api/users", middleware.MiddlewareAuth(handlers.HandlerGetUserByKey))
	app.Post("/v1/api/users", handlers.HandlerCreateUser)
	app.Post("/v1/api/login", handlers.HandlerUserLogin)
	app.Put("/v1/api/users", middleware.MiddlewareAuth(handlers.HandlerUpdateUser))
	app.Delete("/v1/api/users", middleware.MiddlewareAuth(handlers.HandlerDeleteUser))

	app.Get("/v1/api/feeds/:id", handlers.HandlerGetFeedById)
	app.Get("/v1/api/feeds", handlers.HandlerGetFeeds)
	app.Post("/v1/api/feeds", middleware.MiddlewareAuth(handlers.HandlerCreateFeed))
	app.Delete("/v1/api/feeds/:id", middleware.MiddlewareAuth(handlers.HandlerDeleteFeed))

	app.Get("/v1/api/feed-follows", middleware.MiddlewareAuth(handlers.HandlerGetUserFeedFollows))
	app.Post("/v1/api/feed-follows", middleware.MiddlewareAuth(handlers.HandlerCreateFeedFollow))
	app.Delete("/v1/api/feed-follows/:id", middleware.MiddlewareAuth(handlers.HandlerDeleteFeedFollow))

	app.Get("/v1/api/posts", middleware.MiddlewareAuth(handlers.HandlerGetPostsByUser))
	app.Get("/v1/api/posts/:id", handlers.HandlerGetPostsById)

	app.Get("/v1/api/post-likes", middleware.MiddlewareAuth(handlers.HandlerGetPostLikesByUser))
	app.Post("/v1/api/post-likes", middleware.MiddlewareAuth(handlers.HandlerCreatePostLike))
	app.Delete("/v1/api/post_likes/:id", middleware.MiddlewareAuth(handlers.HandlerDeletePostLike))
}
