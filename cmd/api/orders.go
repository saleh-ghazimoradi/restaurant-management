package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/saleh-ghazimoradi/restaurant-management/database"
	"github.com/saleh-ghazimoradi/restaurant-management/models"
	"github.com/saleh-ghazimoradi/restaurant-management/models/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func (app *application) ordersHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := orderCollection.Find(context.TODO(), bson.M{})

	defer cancel()

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var allOrders []bson.M
	if err = result.All(ctx, &allOrders); err != nil {
		log.Fatal(err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"order": allOrders}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) orderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var order models.Order

	err = orderCollection.FindOne(ctx, bson.M{"id": id}).Decode(&order)

	defer cancel()

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"order": order}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createOrderHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var table models.Table
	var order models.Order

	err := app.readJSON(w, r, &order)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if models.ValidatorOrder(*v, &order); !v.Valid() {
		app.serverErrorResponse(w, r, err)
		return
	}

	if order.Table_id != nil {
		err := tableCollection.FindOne(ctx, bson.M{"id": order.Table_id}).Decode(&table)

		defer cancel()

		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	result, insertErr := orderCollection.InsertOne(ctx, order)

	if insertErr != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	defer cancel()

	err = app.writeJSON(w, http.StatusOK, envelope{"order": result}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *application) updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var table models.Table
	var order models.Order

	var updateObj primitive.D

	err = app.readJSON(w, r, &order)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if order.Table_id != nil {

		err := menuCollection.FindOne(ctx, bson.M{"id": order.Table_id}).Decode(&table)

		defer cancel()

		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		updateObj = append(updateObj, bson.E{"menu", order.Table_id})

	}

	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateObj = append(updateObj, bson.E{"updated_at", order.Updated_at})

	upsert := true

	filter := bson.M{"id": id}

	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	result, err := orderCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$st", updateObj},
		},
		&opt,
	)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	defer cancel()

	err = app.writeJSON(w, http.StatusOK, envelope{"order": result}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func OrderItemOrderCreator(order models.Order) string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	order.ID = primitive.NewObjectID()

	order.Order_id = order.ID.Hex()

	orderCollection.InsertOne(ctx, order)
	defer cancel()

	return order.Order_id
}
