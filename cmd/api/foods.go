package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/saleh-ghazimoradi/restaurant-management/models"
	"github.com/saleh-ghazimoradi/restaurant-management/models/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (app *application) foodsHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	recordPerPage, err := strconv.Atoi(r.URL.Query().Get("recordPerPage"))

	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage
	startIndex, err = strconv.Atoi(r.URL.Query().Get("startIndex"))

	matchStage := bson.D{{"$match", bson.D{{}}}}

	groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}

	projectStage := bson.D{
		{
			"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"food_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
			}}}

	result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage})
	defer cancel()

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var allFoods []bson.M

	if err = result.All(ctx, &allFoods); err != nil {
		log.Fatal(err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"foods": allFoods[0]}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) foodHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var food models.Food

	err = foodCollection.FindOne(ctx, bson.M{"id": id}).Decode(&food)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"food": food}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) createFoodHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var menu models.Menu
	var food models.Food

	err := app.readJSON(w, r, &food)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if models.ValidatorFood(*v, &food); !v.Valid() {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = menuCollection.FindOne(ctx, bson.M{"id": food.Menu_id}).Decode(&menu)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.ID = primitive.NewObjectID()
	food.Food_id = food.ID.Hex()

	var num = toFixed(*food.Price, 2)
	food.Price = &num

	result, insertErr := foodCollection.InsertOne(ctx, food)

	if insertErr != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	defer cancel()

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/foods/%d", food.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"food": result}, headers)
	
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateFoodHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var menu models.Menu
	var food models.Food

	err = app.readJSON(w, r, &food)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var updateObj primitive.D

	if food.Name != nil {
		updateObj = append(updateObj, bson.E{"name", food.Name})
	}

	if food.Price != nil {
		updateObj = append(updateObj, bson.E{"price", food.Price})
	}

	if food.Food_image != nil {
		updateObj = append(updateObj, bson.E{"food_image", food.Food_image})
	}

	if food.Menu_id != nil {
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()

		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
		updateObj = append(updateObj, bson.E{"menu", food.Price})

		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		updateObj = append(updateObj, bson.E{"updated_at", food.Updated_at})

		upsert := true

		filter := bson.M{"id": id}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := foodCollection.UpdateOne(
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

		err = app.writeJSON(w, http.StatusOK, envelope{"food": result}, nil)

		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	}
}
