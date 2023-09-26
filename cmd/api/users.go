package main

import (
	"fmt"
	"net/http"
)

func (app *application) usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Show a list of users")
}

func (app *application) userHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
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
