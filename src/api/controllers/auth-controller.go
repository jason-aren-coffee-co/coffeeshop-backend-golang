package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jason-gill00/coffee-shop-backend-golang/src/api/database"

	// "../database"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Testing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("INSIDE THE AUTH CONTROLLER")
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var createUser CreateUser
	err := json.NewDecoder(r.Body).Decode(&createUser)
	if err != nil {
		panic(err)
	}
	log.Print(createUser)
	if createUser.Password == createUser.VerifyPassword {
		user, password := createUser.Username, createUser.Password
		bytePass := []byte(password)
		hashedPassword, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		db := database.Connect()
		database.AddUser(user, string(hashedPassword), db)
		w.Write([]byte(`{"success":true}`))
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	db := database.Connect()
	database.Login(string(user.Username), string(user.Password), db)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    string(user.Username),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	_ = godotenv.Load(".env")
	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	// fmt.Println(token)
	w.Header().Add("auth-token", token)
	auth_token := r.Header.Get("auth-token")
	fmt.Println(auth_token)
	// jwt.Keyfunc()

	decode_claims := jwt.MapClaims{}
	jwt.ParseWithClaims(auth_token, decode_claims, func(decode_token *jwt.Token) (interface{}, error) {
		return []byte(string(os.Getenv("JWT_SECRET"))), nil
	})
	// ans, _ := decode_token.SignedString(string(os.Getenv("JWT_SECRET")))
	// fmt.Println(decode_token.Raw)

	fmt.Println(decode_claims["username"])
	w.Write([]byte(`{"success":true}`))

}
