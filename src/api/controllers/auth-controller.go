package controllers

import (
	"fmt"
	"net/http"
)

func Testing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("INSIDE THE AUTH CONTROLLER")
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Create Account Controller")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Login Controller")
}
