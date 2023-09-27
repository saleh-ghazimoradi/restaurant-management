package main

import (
	"fmt"
	"net/http"
)

func (app *application) orderItemsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "show a list of invoices")
}

func (app *application) orderItemsByOrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "....")
}

func (app *application) orderItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Show an invoice by id: %d", id)
}

func (app *application) createOrderItemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create an invoice")
}

func (app *application) updateOrderItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "update an invoice by id: %d\n", id)
}
