package providers

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/yuricampolongo/crypto-monitoring/user_info_service/src/api/domain"
)

type dynamoDBProvider struct{}

var (
	DynamoDB dynamoDBProvider
	sess     *session.Session
	svc      *dynamodb.DynamoDB
)

func init() {
	DynamoDB = dynamoDBProvider{}
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc = dynamodb.New(sess)
}

func (d *dynamoDBProvider) Save(in interface{}, table string) error {
	data := in.(domain.Entity)

	av, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		log.Fatalf("Got error marshalling map: %s", err)
		return err
	}

	if data.GetId() == "" { //new entity, we create an UUID
		id := uuid.NewString()
		av["id"] = &dynamodb.AttributeValue{
			S: &id,
		}
	}

	// Create item in table tUser
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(table),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
		return err
	}

	fmt.Println("Successfully added item to table " + table)

	return nil
}

func (d *dynamoDBProvider) CreateTable(tableName string) bool {
	if !checkTableExists(tableName) {
		input := &dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       aws.String("HASH"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(10),
				WriteCapacityUnits: aws.Int64(10),
			},
			TableName: aws.String(tableName),
		}

		_, err := svc.CreateTable(input)
		if err != nil {
			log.Fatalf("Got error calling CreateTable: %s", err)
			return false
		}

		fmt.Println("Created the table", tableName)
		return true
	}
	return true // table already exists
}

func checkTableExists(tableName string) bool {
	input := &dynamodb.ListTablesInput{}
	result, err := svc.ListTables(input)
	if err != nil {
		panic(err)
	}
	for _, n := range result.TableNames {
		if *n == tableName {
			return true
		}
	}
	return false
}
