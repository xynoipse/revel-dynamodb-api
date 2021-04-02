package user

import (
	db "revel-dynamodb-api/app/database"
)

// Model is the user model struct
type Model struct {
	UserID string `json:"user_id" dynamo:"UserID"` // PK
	Email  string `json:"email" dynamo:"Email"`    // SK

	FirstName string `json:"first_name" dynamo:"FirstName"`
	LastName  string `json:"last_name" dynamo:"LastName"`
	Password  string `json:"-" dynamo:"Password"`

	db.Model
}
