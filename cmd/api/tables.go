package main

import (
	"fmt"
	"net/http"
)

func (app *application) tablesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "show a list of tables")
}

func (app *application) tableHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Show a table by id: %d", id)
}

func (app *application) createTableHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a table")
}

func (app *application) updateTableHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "update a table by id: %d\n", id)
}
