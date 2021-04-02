package user

import (
	"revel-dynamodb-api/app"
	db "revel-dynamodb-api/app/database"

	"github.com/google/uuid"
)

// Create User item
func (user *Model) Create() (err error) {
	// Set User fields
	user.UserID = uuid.New().String()
	user.PK = db.AppendPrefix(Entity, user.UserID)
	user.SK = db.AppendPrefix(Entity, user.Email)
	user.CreatedAt = db.GetCurrentTimestamp()
	user.UpdatedAt = user.CreatedAt
	user.Entity = Entity

	// Encrypt password
	err = user.EncryptPassword()
	if err != nil {
		return err
	}

	// DynamoDB put operation
	err = app.Table.Put(user).Run()
	if err != nil {
		return
	}

	return nil
}
