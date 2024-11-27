package main

import (
	"fmt"
	"log"
	"os"
	"student-teacher-api/config"
	"student-teacher-api/controller"
	"student-teacher-api/db"
	"student-teacher-api/manager"
	"student-teacher-api/routes"
	"student-teacher-api/service"
	_ "github.com/lib/pq"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// CustomValidator implements Echo's Validator interface
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate validates the struct using go-playground/validator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func main() {
	//config.PostgresConnect()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default configurations.")
	}
	// Connect to the MongoDB
	client, err := db.ConnectDatabase()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Load PostgreSQL configuration
	postgresCfg, err := config.LoadPostgresConfig()
	if err != nil {
		log.Fatalf("Failed to load PostgreSQL configuration: %v", err)
	}
	// Connect to PostgreSQL
	postgresDB, err := db.PostgresConnect(postgresCfg)
	if err != nil {
		log.Fatalf("PostgreSQL connection failed: %v", err)
	}

	// Set the student collection in the service
	service.SetStudentCollection(client, "student")

	// Set the database connection for PostgreSQL
	service.SetDatabase(postgresDB)

	// Initialize the Echo instance
	e := echo.New()

	// Set the default validator to the Echo instance
	e.Validator = &CustomValidator{Validator: validator.New()}

	// Initialize the student service
	studentManager := &manager.StudentManager{}

	// Initialize the controller with the studentService instance
	StudentController := &controller.StudentController{Manager: studentManager}

	// Setup routes with the studentController
	routes.SetupRoutes(e, StudentController)

	// Start the server
	port := os.Getenv("PORT")
	address := fmt.Sprintf(":%s", port)
	if err := e.Start(address); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
