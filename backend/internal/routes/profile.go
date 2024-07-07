package routes

import "backend/internal/handlers"

// Profile implements Routing.
func (r *route) Profile(handler handlers.HandlerImpl) {
	protected := r.echo.Group("/users/:userID/profile")
	protected.GET("", handler.Get)
	protected.PUT("", handler.Create)
}
