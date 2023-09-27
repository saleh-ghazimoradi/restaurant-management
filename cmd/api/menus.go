package main

import (
	"fmt"
	"net/http"
)

func (app *application) menusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "show a list of menus")
}

func (app *application) menuHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Show a menu by id: %d", id)
}

func (app *application) createMenuHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a menu")
}

func (app *application) updateMenuHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "update a menu by id: %d\n", id)
}
