package routes

import "backend/internal/handlers"

// LoginRoute implements Routing.
func (r *route) LoginRoute(handler handlers.Login) {
	r.echo.POST("/auth/login", handler.Login)
}
