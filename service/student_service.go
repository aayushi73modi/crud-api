package service

import (
	"context"
	"fmt"
	"log"
	models "student-teacher-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var studentCollection *mongo.Collection

// SetStudentCollection sets the collection for the service
func SetStudentCollection(client *mongo.Client, database string) {
	studentCollection = client.Database(database).Collection("student")
}

// GetStudents retrieves all Students from the database

func GetStudentsFromMongoDB() ([]models.Student, error) {
	cursor, err := studentCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println("Error fetching students:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var students []models.Student
	for cursor.Next(context.Background()) {
		var student models.Student
		if err := cursor.Decode(&student); err != nil {
			log.Println("Error decoding student:", err)
			continue
		}
		students = append(students, student)
	}
	return students, nil
}

// GetStudentByIDFromMongoDB fetches a Student by ID from MongoDB
func GetStudentByIDFromMongoDB(id primitive.ObjectID) (models.Student, error) {
	var student models.Student
	err := studentCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&student)
	if err != nil {
		log.Println("Error fetching student by ID from MongoDB:", err)
		return models.Student{}, err
	}
	return student, nil
}

// InsertStudent inserts a new Student into the collection
func InsertStudent(student models.Student) (models.Student, error) {
	result, err := studentCollection.InsertOne(context.Background(), student)
	if err != nil {
		return models.Student{}, err
	}
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return models.Student{}, fmt.Errorf("failed to convert InsertedID to ObjectID")
	}

	// Set the ID field in the student struct as a hexadecimal string
	student.ID = objectID.Hex()
	return student, nil
}

// UpdateStudent updates an existing Student
func UpdateStudentInMongoDB(id primitive.ObjectID, student models.Student) (models.Student, error) {
	// Perform the update in MongoDB
	_, err := studentCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": student},
	)
	if err != nil {
		return models.Student{}, err
	}

	// Return the updated student data
	updatedStudent := student
	updatedStudent.ID = id.Hex() // Set the ID as a string
	return updatedStudent, nil
}

// DeleteStudent deletes a Student by its ID
func DeleteStudent(id primitive.ObjectID) error {
	_, err := studentCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error deleting student: %w", err)
	}
	return nil
}
