package db

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() *gorm.DB {

	connStr := os.Getenv("POSTGRES_DB_STRING")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect to postgres: ", err)
	}

	return db
}

func InitializeElasticSearch() *elasticsearch.Client {
	caCert, err := os.ReadFile(os.Getenv("CERT_FILE"))

	if err != nil {
		log.Fatal("Could not read the ca cert file: ", err)
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTIC_DB_STRING"),
		},
		Username: "elastic",
		Password: os.Getenv("ELASTIC_PASSWORD"),
		CACert:   caCert,
	}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatal("Could not connect to elastic search: ", err)
	}

	return es
}
