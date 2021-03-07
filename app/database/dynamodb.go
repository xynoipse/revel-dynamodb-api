package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/revel/revel"
)

// DynamoDB returns DynamoDB client and DynamoDB Table instance
func DynamoDB() (db *dynamo.DB, table dynamo.Table) {
	var (
		tableName string
		region    string
		found     bool
	)

	// AWS config values
	if region, found = revel.Config.String("aws.default.region"); !found {
		revel.RevelLog.Fatal("aws.default.region not configured")
	}
	if tableName, found = revel.Config.String("aws.dynamodb.table"); !found {
		revel.RevelLog.Fatal("aws.dynamodb.table not configured")
	}

	// AWS config
	awsConfig := aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(revel.Config.StringDefault("aws.dynamodb.endpoint", "")),
	}

	// Initialize new session
	session, err := session.NewSession()
	if err != nil {
		revel.AppLog.Fatalf("Failed to initialize a session to AWS: %s", err.Error())
	}

	// Create DynamoDB client
	db = dynamo.New(session, &awsConfig)

	// Create DynamoDB Table instance
	table = db.Table(tableName)

	return
}
