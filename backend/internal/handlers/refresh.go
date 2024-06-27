package handlers

import (
	"backend/internal/config"
	"backend/internal/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Refresh(c echo.Context) error {

	cfg := config.GetInstance()

	type request struct {
		Token string `json:"refreshToken"`
	}

	var refreshToken request

	if err := c.Bind(&refreshToken); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	token, err := jwt.Parse(refreshToken.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Jwt.SecretKey), nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"message": "Failed to parse JWT claims!",
		})
	}

	username, ok := claims["userName"].(string)

	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
			"message": "Invalid Token!",
		})
	}

	password, ok := claims["password"].(string)

	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
			"message": "Invalid Token!",
		})
	}

	accessToken, _, err := utils.GenerateJWT(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"accessToken": accessToken,
	})

}
