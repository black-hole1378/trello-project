package routes

import "backend/internal/handlers"

// UserWorkSpace implements Routing.
func (r *route) UserWorkSpace(handler handlers.HandlerImpl) {
	protected := r.echo.Group("/workspaces/:workspaceId/users")
	protected.GET("", handler.GetAll)
	protected.POST("", handler.Create)
	protected.PUT("/:userID", handler.Update)
	protected.DELETE("/:userID", handler.Delete)
}
