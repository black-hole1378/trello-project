package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
)

// CommentRoute implements Routing.
func (r *route) CommentRoute(handler handlers.CommentHandler) {
	group := r.echo.Group("/tasks/:taskID/comments")
	group.Use(middleware.JWTMiddleware())
	group.POST("", handler.Create)
	group.GET("/:commentID", handler.Get)
}
