package routes

import (
	"backend/internal/handlers"
	"github.com/labstack/echo/v4"
)

func InitRefresh(c *echo.Echo) {
	c.POST("api/refresh", handlers.Refresh)
}
