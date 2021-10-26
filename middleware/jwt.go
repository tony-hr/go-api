package middleware

import (
	"go-api/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string) (string, error) {
	conf := config.GetConfig()
	jwtKey := conf.GetString("jwt.secret_key")
	jwtDuration := conf.GetInt("jwt.duration")

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(jwtDuration)).Unix()

	t, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return t, err
	}

	return t, err
}
