package models

import (
	"time"

	"github.com/saleh-ghazimoradi/restaurant-management/models/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name"`
	Last_name     *string            `json:"last_name"`
	Password      *string            `json:"password"`
	Email         *string            `json:"email"`
	Avatar        *string            `json:"avatar"`
	Phone         *string            `json:"phone"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}

func ValidatorUser(v validator.Validator, user *User) {
	v.Check(*user.First_name != "", "First Name", "first name is required")
	v.Check(len(*user.First_name) >= 2 && len(*user.First_name) <= 100, "First Name", "First Name must be in the range of 2 to 100 bytes long")
	v.Check(*user.Last_name != "", "Last Name", "last name is required")
	v.Check(len(*user.Last_name) >= 2 && len(*user.Last_name) <= 100, "Last Name", "Last Name must be in the range of 2 to 100 bytes long")
	v.Check(*user.Password != "", "password", "password is required")
	v.Check(len(*user.Password) >= 6 && len(*user.Password) <= 20, "password", "password must be in the range of 6 to 20 bytes long")
	v.Check(*user.Email != "", "email", "must be provided")
	v.Check(*user.Phone != "", "phone", "must be provided")
}
