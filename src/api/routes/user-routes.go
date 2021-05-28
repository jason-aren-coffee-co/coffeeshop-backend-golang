package routes

import (
	"net/http"

	"github.com/jason-gill00/coffee-shop-backend-golang/src/api/controllers"
)

var authRoutes = []Route{
	{
		Url:     "/api/auth/signup",
		Method:  http.MethodPost,
		Handler: controllers.CreateAccount,
	},
	{
		Url:     "/api/auth/login",
		Method:  http.MethodPost,
		Handler: controllers.Login,
	},
}
