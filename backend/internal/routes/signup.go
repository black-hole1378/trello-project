package routes

import "backend/internal/handlers"

// SignUp implements Routing.
func (r *route) SignUp(handler handlers.HandlerImpl) {
	r.echo.POST("/auth/signup", handler.Create)
}
