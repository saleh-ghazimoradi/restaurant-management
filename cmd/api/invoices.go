package main

import (
	"fmt"
	"net/http"
)

func (app *application) invoicesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "show a list of invoices")
}

func (app *application) invoiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Show an invoice by id: %d", id)
}

func (app *application) createInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create an invoice")
}

func (app *application) updateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "update an invoice by id: %d\n", id)
}
