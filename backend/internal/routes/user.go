package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
)

func (r *route) UserRoute(handler handlers.HandlerImpl) {
	r.echo.POST("/users", handler.Create)
	r.echo.GET("/users", handler.GetAll)
	userGroup := r.echo.Group("/users")
	userGroup.Use(middleware.JWTMiddleware())
	userGroup.GET("/:userID", handler.Get)
	userGroup.PUT("/:userID", handler.Update)
	userGroup.DELETE("/:userID", handler.Delete)
}
