package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func VerifyToken(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "IN THE MIDDLEWARE")
		_ = godotenv.Load(".env")

		if r.Header["Authtoken"] != nil {
			// fmt.Println("header token: ", r.Header["Authtoken"])
			token, err := jwt.Parse(r.Header["Authtoken"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				panic(err)
			}
			if token.Valid {
				fmt.Println("VALID")
				handler.ServeHTTP(w, r)
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Not Authorized")
		}

		// handler.ServeHTTP(w, r)
	}
}
