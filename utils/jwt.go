package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenToken(email string) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)

	claims := t.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["expires"] = time.Now().Add(time.Minute * 20).Unix()

	token, err := t.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return token, nil

}
