package routes

import (
	"net/http"

	// "../middleware"
	"github.com/jason-gill00/coffee-shop-backend-golang/src/api/controllers"
	"github.com/jason-gill00/coffee-shop-backend-golang/src/api/middleware"
)

var orderRoutes = []Route{
	{
		Url:     "/api/order/test",
		Method:  http.MethodGet,
		Handler: controllers.TestingFunc,
	},
	{
		Url:     "/api/order/history",
		Method:  http.MethodGet,
		Handler: middleware.VerifyToken(controllers.GetHistory),
	},
	{
		Url:     "/api/order/submit",
		Method:  http.MethodPost,
		Handler: controllers.CreateOrder,
	},
}
