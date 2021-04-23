package db

import (
	"log"
	"os"

	"github.com/yuricampolongo/crypto-monitoring/user_info_service/src/api/db/providers"
)

type dbInterface interface {
	CreateTable(tableName string) bool
	Save(in interface{}, table string) error
}

func Provider() dbInterface {
	provider, err := os.LookupEnv("CRYPTO_MONITORING_DB_PROVIDER")
	if err {
		log.Println("No CRYPTO_MONITORING_DB_PROVIDER env variable set, DynamoDB will be used")
		provider = "dynamodb"
	}

	switch provider {
	default:
		return &providers.DynamoDB
	}

}
