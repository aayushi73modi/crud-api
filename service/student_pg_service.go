package service

import (
	"fmt"
	"log"
	models "student-teacher-api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB

// SetDatabase initializes the GORM database connection
func SetDatabase(database *gorm.DB) {
	db = database
}

func InsertStudentps(student models.Student) (models.Student, error) {
	// Before inserting, ensure the table exists (AutoMigrate is already called in PostgresConnect)
	err := db.AutoMigrate(&models.Student{}) // Create the table if it doesn't exist
	if err != nil {
		log.Println("Error during AutoMigrate:", err)
		return models.Student{}, fmt.Errorf("error ensuring table exists: %w", err)
	}

	// Now insert the student record
	err = db.Create(&student).Error
	if err != nil {
		log.Println("Error inserting student:", err)
		return models.Student{}, fmt.Errorf("error inserting student: %w", err)
	}

	return student, nil
}

// GetStudents retrieves all Students from the database
func GetStudentsFromPostgreSQL() ([]models.Student, error) {
	var students []models.Student
	err := db.Find(&students).Error
	if err != nil {
		log.Println("Error fetching students:", err)
		return nil, err
	}
	return students, nil
}

// GetStudentByIDFromPostgreSQL fetches a Student by ID from PostgreSQL
func GetStudentByIDFromPostgreSQL(id string) (models.Student, error) {
	var student models.Student
	err := db.Where("id = ?", id).First(&student).Error
	if err != nil {
		log.Println("Error fetching student by ID from PostgreSQL:", err)
		return models.Student{}, err
	}
	return student, nil
}

// UpdateStudent updates an existing Student
func UpdateStudentInPostgreSQL(id string, student models.Student) (models.Student, error) {
	// Perform the update in PostgreSQL
	err := db.Model(&models.Student{}).Where("id = ?", id).Updates(student).Error
	if err != nil {
		return models.Student{}, fmt.Errorf("error updating student: %w", err)
	}

	// Return the updated student data
	student.ID = id // Set the ID (PostgreSQL uses integers for IDs)
	return student, nil
}

// DeleteStudent deletes a Student by its ID
func DeleteStudents(id string) error {
	studentID, err := uuid.Parse(id)
	if err != nil {
		log.Println("Error parsing UUID:", err)
		return fmt.Errorf("invalid UUID format: %w", err)
	}
	//delete the student by UUID
	err = db.Delete(&models.Student{}, "id = ?", studentID).Error
	if err != nil {
		log.Println("Error deleting student:", err)
		return fmt.Errorf("error deleting student: %w", err)
	}
	return nil
}
