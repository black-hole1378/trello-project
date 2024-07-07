package handlers

import (
	"backend/internal/repositories"
	"backend/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Login interface {
	Login(c echo.Context) error
}

type login struct {
	repository repositories.UserRepoImpl
}

// Login implements Login.
func (l *login) Login(c echo.Context) error {
	var req struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request payload",
		})
	}

	user, err := l.repository.GetUserByUserName(req.UserName)

	if err != nil || user == nil {
		return handleDBError(err, c)
	}

	if !utils.ComparePasswordHash(req.Password, user.Password) {
		return handleDBError(gorm.ErrRecordNotFound, c)
	}

	accessToken, refreshToken, err := utils.GenerateJWT(req.UserName, req.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"userName":     user.Username,
		"userID":       user.ID,
		"message":      "Welcome back!",
	})
}

func NewLoginHandler() Login {
	return &login{
		repository: repositories.NewUserRepository(),
	}
}
