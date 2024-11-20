package service

import (
	"database/sql"
	"fmt"
	"log"
	models "student-teacher-api/model"

	_ "github.com/lib/pq" // PostgreSQL driver
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db *sql.DB

// SetDatabase initializes the database connection
func SetDatabase(database *sql.DB) {
	db = database

}

// GetStudents retrieves all Students from the database
func GetStudentsFromPostgreSQL() ([]models.Student, error) {
	rows, err := db.Query("SELECT id, name, age, class FROM studentsps")
	if err != nil {
		log.Println("Error fetching students:", err)
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.Student_name, &student.Age, &student.Class); err != nil {
			log.Println("Error scanning student:", err)
			continue
		}
		students = append(students, student)
	}
	return students, nil
}

// GetStudentByID fetches a Student by its ID
func GetStudentsByID(id int) (models.Student, error) {
	var student models.Student
	err := db.QueryRow("SELECT id, name, age, class FROM students WHERE id = $1", id).Scan(
		&student.ID, &student.Student_name, &student.Age, &student.Class)
	if err != nil {
		return models.Student{}, err
	}
	return student, nil
}

// InsertStudent inserts a new Student into the database
func InsertStudentps(student models.Student) (models.Student, error) {
	var id int
	err := db.QueryRow(
		"INSERT INTO students (id,name, age, class) VALUES ($1, $2, $3,$4) RETURNING id",
		student.ID, student.Student_name, student.Age, student.Class,
	).Scan(&id)
	log.Println("Insert record in postgres.")
	if err != nil {
		return models.Student{}, err
	}
	//student.ID = result.InsertedID.(primitive.ObjectID)
	student.ID = primitive.NewObjectID()
	return student, nil
}

// UpdateStudent updates an existing Student
func UpdateStudents(id int, student models.Student) error {
	_, err := db.Exec(
		"UPDATE students SET name = $1, age = $2, class = $3 WHERE id = $4",
		student.ID, student.Student_name, student.Age, student.Class, id,
	)
	if err != nil {
		return fmt.Errorf("error updating student: %w", err)
	}
	return nil
}

// DeleteStudent deletes a Student by its ID
func DeleteStudents(id int) error {
	_, err := db.Exec("DELETE FROM students WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting student: %w", err)
	}
	return nil
}
