package server

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/routes"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	// database instance creating.
	_ = database.GetInstance()

	// echo instance
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// config instance
	cfg := config.GetInstance()

	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	// init routes
	routing := routes.NewRoute(e)
	routing.UserRoute(handlers.NewUserHandler())
	routing.WorkSpaceRoute(handlers.NewWorkSpaceHandler())
	routing.TaskRoute(handlers.NewTaskHandler())
	routing.LoginRoute(handlers.NewLoginHandler())
	routing.SubTaskRoute(handlers.NewSubTaskHandler())
	routes.InitRefresh(e)
	routing.SignUp(handlers.NewUserHandler())
	routing.Profile(handlers.NewUserHandler())
	routing.UserWorkSpace(handlers.NewUserWorkSpaceHandler())
	routing.CommentRoute(*handlers.NewCommentHandler())
	// starting server.
	e.Logger.Fatal(e.Start(address))
}
