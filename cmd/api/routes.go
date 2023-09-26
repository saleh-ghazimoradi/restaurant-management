package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/users", app.usersHandler)

	router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.userHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users/signup", app.userSignupHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users/login", app.userLoginHandler)

	router.HandlerFunc(http.MethodGet, "/v1/foods", app.foodsHandler)

	router.HandlerFunc(http.MethodGet, "/v1/foods/:id", app.foodHandler)

	router.HandlerFunc(http.MethodPost, "/v1/foods", app.createFoodHandler)

	router.HandlerFunc(http.MethodPatch, "/v1/foods/:id", app.updateFoodHandler)

	return router

}
