package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type Message struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
}

func IsAuthorizationToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		autorization := r.Header.Get("Authorization")
		log.Println("new req ", autorization)
		if len(autorization) > 0 {

			authToken := strings.Split(autorization, " ")

			token, err := jwt.Parse(authToken[1], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, fmt.Errorf("error parsing")
				}

				return []byte(os.Getenv("JWT_KEY")), nil

			})

			if err != nil {
				if errors.Is(err, jwt.ErrTokenExpired) {
					var message Message
					message.Error = true
					message.Msg = "Session expired"

					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)

					json.NewEncoder(w).Encode(message)
					return
				}

			}

			_, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				next.ServeHTTP(w, r)
			}
		} else {
			var message Message
			message.Error = true
			message.Msg = "Forbidden"

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)

			json.NewEncoder(w).Encode(message)

			return
		}

	})
}

func Cors(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		if origin != "" {

			if origin == "http://localhost:3000" {

				w.Header().Add("Vary", "Origin")
				w.Header().Add("Vary", "Access-Control-Request-Method")

				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
				w.Header().Add("Access-Control-Allow-Credentials", "true")

				if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
					w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, PATCH, PUT, DELETE, GET")
					w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

					w.WriteHeader(http.StatusOK)
					return
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
