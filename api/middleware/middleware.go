package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func IsAuthorizationToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		autorization := r.Header.Get("Authorization")

		authToken := strings.Split(autorization, " ")

		token, err := jwt.Parse(authToken[1], func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Some error")
			}

			return []byte(os.Getenv("JWT_KEY")), nil

		})

		if err != nil {

			var Mesages struct {
				Error bool   `json:"error"`
				Msg   string `json:"msg"`
			}

			if errors.Is(err, jwt.ErrTokenExpired) {
				Mesages.Error = true
				Mesages.Msg = "Session expired"
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				json.NewEncoder(w).Encode(Mesages)
				return
			}

		}

		_, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			next.ServeHTTP(w, r)
		}

	})
}
