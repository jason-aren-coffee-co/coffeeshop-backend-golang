package routes

import (
	"net/http"

	"github.com/jason-gill00/coffee-shop-backend-golang/src/api/controllers"
)

var orderRoutes = []Route{
	{
		Url:     "/api/order/history",
		Method:  http.MethodGet,
		Handler: controllers.GetHistory,
	},
	{
		Url:     "/api/order/submit",
		Method:  http.MethodPost,
		Handler: controllers.CreateOrder,
	},
}
