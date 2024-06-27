package routes

import (
	"backend/internal/handlers"

	"github.com/labstack/echo/v4"
)

type Routing interface {
	WorkSpaceRoute(handler handlers.HandlerImpl)
	UserRoute(handler handlers.HandlerImpl)
	TaskRoute(handler handlers.HandlerImpl)
	LoginRoute(handler handlers.Login)
	SubTaskRoute(handler handlers.HandlerImpl)
}

type route struct {
	echo *echo.Echo
}

func NewRoute(echo *echo.Echo) Routing {
	return &route{
		echo: echo,
	}
}
