package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"student-teacher-api/config"
	models "student-teacher-api/model"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *mongo.Client
var PG *gorm.DB

// Load environment variables

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func PostgresConnect(cfg *config.PostgresConfig) (*gorm.DB, error) {
	// Construct the DSN
	postgresDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.POSTGRES_HOST,
		cfg.POSTGRES_PORT,
		cfg.POSTGRES_USER,
		cfg.POSTGRES_PASSWORD,
		cfg.POSTGRES_DB,
	)
	log.Printf("Postgres DSN: %s", postgresDSN)

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to PostgreSQL: %w", err)
	}
	log.Println("Connected to PostgreSQL successfully!")
	PG = db

	// Automatically migrate the database schema for the Student model
	if err := db.AutoMigrate(&models.Student{}); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate: %w", err)
	}

	return db, nil
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
