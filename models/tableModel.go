package models

import (
	"time"

	"github.com/saleh-ghazimoradi/restaurant-management/models/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID               primitive.ObjectID `bson:"_id"`
	Number_of_guests *int               `json:"number_of_guests"`
	Table_number     *int               `json:"table_number"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
	Table_id         string             `json:"table_id"`
}

func ValidatorTable(v validator.Validator, table *Table) {
	v.Check(*table.Number_of_guests != 0, "Number of Guests", "Number of guests must be provided")
	v.Check(*table.Table_number != 0, "Table number", "table number is required")
}
