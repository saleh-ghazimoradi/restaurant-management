package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/users", app.usersHandler)

	router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.userHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users/signup", app.userSignupHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users/login", app.userLoginHandler)

	router.HandlerFunc(http.MethodGet, "/v1/foods", app.foodsHandler)

	router.HandlerFunc(http.MethodGet, "/v1/foods/:id", app.foodHandler)

	router.HandlerFunc(http.MethodPost, "/v1/foods", app.createFoodHandler)

	router.HandlerFunc(http.MethodPatch, "/v1/foods/:id", app.updateFoodHandler)

	router.HandlerFunc(http.MethodGet, "/v1/invoices", app.invoicesHandler)

	router.HandlerFunc(http.MethodGet, "/v1/invoices/:invoice_id", app.invoiceHandler)

	router.HandlerFunc(http.MethodPost, "/v1/invoices", app.createInvoiceHandler)

	router.HandlerFunc(http.MethodPatch, "/v1/invoices/:invoice_id", app.updateInvoiceHandler)

	router.HandlerFunc(http.MethodGet, "/v1/menus", app.menusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/menus/:menu_id", app.menuHandler)

	router.HandlerFunc(http.MethodPost, "/v1/menus", app.createMenuHandler)

	router.HandlerFunc(http.MethodPatch, "/v1/menu/:menu_id", app.updateMenuHandler)

	router.HandlerFunc(http.MethodGet, "/v1/orders", app.ordersHandler)

	router.HandlerFunc(http.MethodGet, "/v1/orders/:order_id", app.orderHandler)

	router.HandlerFunc(http.MethodPost, "/v1/orders", app.createOrderHandler)

	router.HandlerFunc(http.MethodPatch, "/v1/orders/:order_id", app.updateOrderHandler)

	router.HandlerFunc(http.MethodGet, "/v1/tables", app.tablesHandler)

	router.HandlerFunc(http.MethodGet, "/v1/tables/:table_id", app.tableHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tables", app.createTableHandler)

	router.HandlerFunc(http.MethodPatch, "/v1/tables/:table_id", app.updateTableHandler)

	router.HandlerFunc(http.MethodGet, "/v1/orderItems", app.orderItemsHandler)

	router.HandlerFunc(http.MethodGet, "/v1/orderItems/:orderItem_id", app.orderItemHandler)

	router.HandlerFunc(http.MethodGet, "/v1/orderItems-order/:order_id", app.orderItemsByOrderHandler)

	router.HandlerFunc(http.MethodPost, "/v1/orderItems", app.createOrderItemHandler)

	router.HandlerFunc(http.MethodPatch, "/v1/orderItems/:orderItem_id", app.updateOrderItemHandler)

	return router

}
