package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
)

func (r *route) TaskRoute(handler handlers.HandlerImpl) {
	taskGroup := r.echo.Group("/workspaces/:workspaceId/tasks")
	taskGroup.Use(middleware.JWTMiddleware())
	taskGroup.GET("", handler.GetAll)
	taskGroup.POST("", handler.Create)
	taskGroup.GET("/:taskID", handler.Get)
	taskGroup.PUT("/:taskID", handler.Update)
	taskGroup.DELETE("/:taskID", handler.Delete)
}
