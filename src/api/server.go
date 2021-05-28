package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jason-gill00/coffee-shop-backend-golang/src/api/routes"
)

func Run() {
	r := mux.NewRouter()
	routes.SetupRoutes(r)
	log.Fatal(http.ListenAndServe(":5000", r))

}
