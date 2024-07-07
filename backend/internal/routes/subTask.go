package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
)

// SubTaskRoute implements Routing.
func (r *route) SubTaskRoute(handler handlers.HandlerImpl) {
	taskGroup := r.echo.Group("/tasks/:taskID/subtasks")
	taskGroup.Use(middleware.JWTMiddleware())
	taskGroup.GET("", handler.GetAll)
	taskGroup.POST("", handler.Create)
	taskGroup.GET("/:subtaskID", handler.Get)
	taskGroup.PUT("/:subtaskID", handler.Update)
	taskGroup.DELETE("/:subtaskID", handler.Delete)
}
