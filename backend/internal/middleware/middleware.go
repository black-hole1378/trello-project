package middleware

import (
	"backend/internal/config"
	"backend/internal/repositories"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cfg := config.GetInstance()

			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Invalid Token Please try Again1!",
				})
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Invalid Token Please try Again2!",
				})
			}

			tokenString := tokenParts[1]

			// Parse and verify the JWT token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Use the secret key from the configuration
				return []byte(cfg.Jwt.SecretKey), nil
			})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Invalid Token Please try Again3!",
				})
			}

			// Check if the token is valid
			if !token.Valid {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Invalid Token Please try Again4!",
				})
			}

			// Extract user information from the token and set it in the request context
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Unable to parse token!",
				})
			}

			userName, ok := claims["userName"].(string)

			if !ok {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Unable to parse token",
				})
			}

			password, ok := claims["password"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Invalid Token Please try Again5!",
				})
			}

			fmt.Println(userName)

			repo := repositories.NewUserRepository()

			if repo.Check(userName, password) == nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Invalid Token Please try Again6!",
				})
			}

			c.Set("userName", userName)
			c.Set("password", password)

			return next(c)
		}
	}
}
