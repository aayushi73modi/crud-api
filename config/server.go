package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var PG *sql.DB

// Load environment variables
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func PostgresConnect() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")

	postgresDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", postgresHost, postgresPort, postgresUser, postgresPassword, postgresDB)
	PGDB, err := sql.Open("postgres", postgresDSN)
	if err != nil {
		log.Fatal("Error connecting to PostgreSQL:", err)
	}

	err = PGDB.Ping()
	if err != nil {
		log.Fatal("Error pinging PostgreSQL:", err)
	}
	if PGDB == nil {
		log.Println("Database connection is nil!")
	}
	PG = PGDB
	log.Println("Connected to PostgreSQL successfully!")
	return PGDB, nil
}

// ConnectDatabase initializes the MongoDB client
func ConnectDatabase() (*mongo.Client, error) {
	LoadEnv()
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("could not connect to MongoDB: %w", err)
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}
