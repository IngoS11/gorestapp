package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := assembleDatabaseDSN()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to db, error: %v\n", err)
	}
}

// assembleDatabaseDNS godoc
// assembles the datbase dsn from environment variables.
func assembleDatabaseDSN() string {

	dbhost := os.Getenv("POSTGRES_HOSTNAME")
	if dbhost == "" {
		log.Fatal("Missing database hostname")
	}

	dbport := os.Getenv("POSTGRES_PORT")
	if dbport == "" {
		log.Println("Missing environment variable POSTGRES_PORT, using default port: 5432")
	} else {
		dbport = "5432"
	}

	db := os.Getenv("POSTGRES_DB")
	if db == "" {
		log.Fatal("Missing environment variable POSTGRES_DB")
	}

	dbuser := os.Getenv("POSTGRES_USER")
	if dbuser == "" {
		log.Fatal("Missing environment variable POSTGRES_USER")
	}

	dbpw := os.Getenv("POSTGRES_PASSWORD")
	if dbpw == "" {
		log.Fatal("Missing environment variable POSTGRES_PASSWORD")
	}

	dbattr := os.Getenv("POSTGRES_ATTRIBUTES")

	connectionDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s",
		dbhost, dbuser, dbpw, db, dbport, dbattr)

	return connectionDSN
}
