package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Show a list of users")
}

func (app *application) userHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Show a user by id: %d", id)
}

func (app *application) userSignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "userSigned up")
}

func (app *application) userLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user logged in")
}
