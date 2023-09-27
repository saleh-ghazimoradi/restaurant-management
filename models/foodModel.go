package models

import (
	"time"

	"github.com/saleh-ghazimoradi/restaurant-management/models/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name"`
	Price      *float64           `json:"price"`
	Food_image *string            `json:"food_image"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Food_id    string             `json:"food_id"`
	Menu_id    *string            `json:"menu_id"`
}

func ValidatorFood(v validator.Validator, food *Food) {
	v.Check(*food.Name != "", "name", "must be provided")
	v.Check(len(*food.Name) >= 2, "", "must be at least 2 bytes long")
	v.Check(len(*food.Name) <= 100, "", "must not be more than 100 bytes long")

	v.Check(*food.Food_image != "", "food_image", "food image must be provided")
	v.Check(*food.Menu_id != "", "food_menu", "menu id must be provided")
}
