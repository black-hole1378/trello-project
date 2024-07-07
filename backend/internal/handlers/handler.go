package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HandlerImpl interface {
	Create(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
	GetAll(c echo.Context) error
	Get(c echo.Context) error
}

func handleDBError(err error, c echo.Context) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": err.Error(),
		})
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{
		"message": err.Error(),
	})
}
