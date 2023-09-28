package main

import (
	"context"
	"fmt"
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

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func (app *application) menusHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := menuCollection.Find(context.TODO(), bson.M{})
	defer cancel()

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var allMenus []bson.M

	if err = result.All(ctx, &allMenus); err != nil {
		log.Fatal(err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"menu": allMenus}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) menuHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show a menu id %d\n", id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var menu models.Menu

	err = foodCollection.FindOne(ctx, bson.M{"id": id}).Decode(&menu)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"menu": menu}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) createMenuHandler(w http.ResponseWriter, r *http.Request) {
	var menu models.Menu
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := app.readJSON(w, r, &menu)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if models.ValidatorMenu(*v, &menu); !v.Valid() {
		app.serverErrorResponse(w, r, err)
		return
	}

	menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	menu.ID = primitive.NewObjectID()
	menu.Menu_id = menu.ID.Hex()

	result, insertErr := menuCollection.InsertOne(ctx, menu)
	if insertErr != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/menus/%d", menu.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"menu": result}, headers)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func inTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}

func (app *application) updateMenuHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var menu models.Menu

	err = app.readJSON(w, r, &menu)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	filter := bson.M{"id": id}

	var updateObj primitive.D

	if menu.Start_Date != nil && menu.End_Date != nil {
		if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
			app.serverErrorResponse(w, r, err)
			return
		}

		updateObj = append(updateObj, bson.E{"start_date", menu.Start_Date})
		updateObj = append(updateObj, bson.E{"end_date", menu.End_Date})

		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{"name", menu.Name})
		}

		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{"name", menu.Category})
		}

		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})

		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := menuCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		err = app.writeJSON(w, http.StatusCreated, envelope{"menu": result}, nil)

		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	}
}
