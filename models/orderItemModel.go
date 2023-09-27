package models

import (
	"time"

	"github.com/saleh-ghazimoradi/restaurant-management/models/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	Quantity      *string            `json:"quantity"`
	Unit_price    *float64           `json:"unit_price"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	Food_id       *string            `json:"food_id"`
	Order_item_id string             `json:"order_item_id"`
	Order_id      string             `json:"order_id"`
}

func ValidatorOrderItem(v validator.Validator, orderItem *OrderItem) {
	v.Check(*orderItem.Quantity == "S" || *orderItem.Quantity == "M" || *orderItem.Quantity == "L", "Quantity", "Quantity must be one of these letters S, M, L")
	v.Check(*orderItem.Unit_price != 0, "Unit_price", "Unit price must be provided")
	v.Check(*orderItem.Food_id != "", "food_id", "food_id must be provided")
	v.Check(orderItem.Order_id != "", "order_id", "order_id must be provided")
}
