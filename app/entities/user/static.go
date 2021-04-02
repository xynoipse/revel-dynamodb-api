package user

import (
	"revel-dynamodb-api/app"
	db "revel-dynamodb-api/app/database"

	"github.com/guregu/dynamo"
)

// FindByEmail return user with the specified email address
func FindByEmail(email string) (result Model, err error) {
	// Query items with SK email address and PK begins with User Type
	err = app.Table.Get(db.SortKey, db.AppendPrefix(Entity, email)).
		Range(db.PartitionKey, dynamo.BeginsWith, Entity).
		Index(db.InvertedIndex).
		One(&result)

	if err != nil {
		return
	}

	return
}
