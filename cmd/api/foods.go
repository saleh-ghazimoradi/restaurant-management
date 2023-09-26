package main

import (
	"fmt"
	"net/http"
)

func (app *application) foodsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "show a list of food")
}

func (app *application) foodHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of food %d\n", id)
}

func (app *application) createFoodHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a dish of food")
}

func (app *application) updateFoodHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "update a dish of food %d\n", id)
}
