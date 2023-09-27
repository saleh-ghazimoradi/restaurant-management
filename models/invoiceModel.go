package models

import (
	"time"

	"github.com/saleh-ghazimoradi/restaurant-management/models/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string             `json:"invoice_id`
	Order_id         string             `json:"order_id"`
	Payment_method   *string            `json:"payment_method"`
	Payment_status   *string            `json:"payment_status"`
	Payment_due_date time.Time          `json:"payment_due_date"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
}

func ValidatorInvoice(v validator.Validator, invoice *Invoice) {
	v.Check(*invoice.Payment_method == "CARD" || *invoice.Payment_method == "CASH" || *invoice.Payment_method == "", "payment_method", "must be one of these methods")
	v.Check(*invoice.Payment_status != "", "payment_status", "payment status is required")
	v.Check(*invoice.Payment_status == "PENDING" || *invoice.Payment_status == "PAID", "payment_status", "payment status must be assigned")
} 
