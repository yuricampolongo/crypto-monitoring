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

type dynamoDBInterface interface {
	Save(user domain.User) error
}
type dynamoDB struct{}

const (
	userTable = "tUser"
)

var (
	DynamoDB dynamoDBInterface
	sess     *session.Session
	svc      *dynamodb.DynamoDB
)

func init() {
	DynamoDB = &dynamoDB{}
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc = dynamodb.New(sess)
	createUserTable()
}

func (d *dynamoDB) Save(user domain.User) error {
	if user.Id == nil {
		//new user, we create an UUID
		uuid := uuid.NewString()
		user.Id = &uuid
	}

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Fatalf("Got error marshalling map: %s", err)
		return err
	}

	// Create item in table tUser
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(userTable),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
		return err
	}

	fmt.Println("Successfully added '" + user.Name + "' (" + (*user.Id) + ") to table " + userTable)

	return nil
}

func createUserTable() {
	if !checkTableExists() {
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
			TableName: aws.String(userTable),
		}

		_, err := svc.CreateTable(input)
		if err != nil {
			log.Fatalf("Got error calling CreateTable: %s", err)
		}

		fmt.Println("Created the table", userTable)
	}
}

func checkTableExists() bool {
	input := &dynamodb.ListTablesInput{}
	result, err := svc.ListTables(input)
	if err != nil {
		panic(err)
	}
	for _, n := range result.TableNames {
		if *n == userTable {
			return true
		}
	}
	return false
}
