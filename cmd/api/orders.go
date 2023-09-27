package main

import (
	"fmt"
	"net/http"
)

func (app *application) ordersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "show a list of orders")
}

func (app *application) orderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Show an order by id: %d", id)
}

func (app *application) createOrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create an order")
}

func (app *application) updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "update an order by id: %d\n", id)
}
