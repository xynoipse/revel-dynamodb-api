package database

// Model struct
type Model struct {
	PK string `json:"-" dynamo:"PK"` // Hash key, a.k.a. partition key
	SK string `json:"-" dynamo:"SK"` // Range key, a.k.a. sort key

	CreatedAt int64  `json:"created_at" dynamo:"CreatedAt"`
	UpdatedAt int64  `json:"updated_at" dynamo:"UpdatedAt"`
	Entity    string `json:"-" dynamo:"Entity"`
}
