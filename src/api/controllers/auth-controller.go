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
		panic(err)
	}
	fmt.Println(token)
	w.Header().Add("auth-token", token)

}
