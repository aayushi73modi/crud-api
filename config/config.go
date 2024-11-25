package config

import (
	"context"
	"fmt"
	"log"
	"os"
	models "student-teacher-api/model"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *mongo.Client
var PG *gorm.DB

type MongoConfig struct {
	MongoDBURL           string `env:"MONGODB_URL" envDefault:"mongodb://localhost:27017"`
	MongoDBName          string `env:"MONGODB_DB_NAME" envDefault:"student"`
	SetStudentCollection string `env:"SetStudentCollection" envDefault:"student"`
}

type PostgresConfi struct {
	POSTGRES_HOST string `env:"POSTGRES_HOST" envDefault:"localhost"`
	POSTGRES_PORT string `env:"POSTGRES_PORT" envDefault:"5433"`
	//POSTGRES_USER     string `env:"POSTGRES_USER" envDefault:""root"`
	POSTGRES_PASSWORD string `env:"POSTGRES_PASSWORD" envDefault:"root"`
	POSTGRES_DB       string `env:"POSTGRES_DB" envDefault:"studentpostgres"`
	//SetStudentCollection string `env:"SetStudentCollection" envDefault:"student"`
}

// Load environment variables
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func PostgresConnect() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to PostgreSQL:", err)
	}

	PG = db
	log.Println("Connected to PostgreSQL successfully!")
	err = db.AutoMigrate(&models.Student{})
	if err != nil {
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
