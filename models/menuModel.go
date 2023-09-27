package models

import (
	"time"

	"github.com/saleh-ghazimoradi/restaurant-management/models/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name"`
	Category   string             `json:"category"`
	Start_Date *time.Time         `json:"start_date"`
	End_Date   *time.Time         `json:"end_date"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Menu_id    string             `json:"menu_id"`
}

func Validator(v validator.Validator, menu *Menu) {
	v.Check(menu.Name != "", "name", "must be provided")
	v.Check(menu.Category != "", "category", "must be provided")
}
