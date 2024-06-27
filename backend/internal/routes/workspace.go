package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
)

func (r *route) WorkSpaceRoute(handler handlers.HandlerImpl) {
	protected := r.echo.Group("/workspaces")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("", handler.GetAll)
	protected.POST("", handler.Create)
	protected.GET("/:workspaceId", handler.Get)
	protected.PUT("/:workspaceId", handler.Update)
	protected.DELETE("/:workspaceId", handler.Delete)
}
