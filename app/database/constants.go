package database

const (
	// Table primary keys
	PartitionKey = "PK" // Hash key
	SortKey      = "SK" // Range key

	// Global Secondary Indexes
	InvertedIndex = "InvertedIndex"
	GetIndexItems = "GetIndexItems"
)
