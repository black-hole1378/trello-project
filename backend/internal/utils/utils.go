package utils

import (
	"backend/internal/config"
	"errors"
	"github.com/dlclark/regexp2"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func IsValidEmail(email string) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if !checkRegex(regex, email) {
		return errors.New("invalid Email")
	}

	return nil

}

func checkRegex(regex string, value string) bool {
	re, _ := regexp2.Compile(regex, regexp2.None)

	match, err := re.MatchString(value)

	if err != nil {
		log.Info(err.Error())
	}

	return match
}

func PasswordCheck(password string) error {
	regex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[^\da-zA-Z]).{8,}$`

	if !checkRegex(regex, password) {
		return errors.New("invalid Password")
	}

	return nil
}

func GenerateJWT(userName string, password string) (string, string, error) {
	cfg := config.GetInstance()
	claims := jwt.MapClaims{
		"userName": userName,
		"password": password,
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString([]byte(cfg.Jwt.SecretKey))

	if err != nil {
		return "", "", err
	}

	claims = jwt.MapClaims{
		"userName": userName,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.Jwt.SecretKey))

	if err != nil {
		return "", "", nil
	}

	return accessTokenString, refreshTokenString, nil
}

func HashedPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Info(err.Error())
	}

	return string(bytes)
}

func ComparePasswordHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
