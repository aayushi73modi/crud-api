package manager

import (
	"fmt"
	"student-teacher-api/Request"
	"student-teacher-api/Response"
	models "student-teacher-api/model"
	"student-teacher-api/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StudentManager handles business logic for Students
type StudentManager struct{}

// GetStudents fetches all Students based on the data source
func (m *StudentManager) GetStudents(flag bool) ([]Response.StudentResponse, error) {
	var students []models.Student
	var err error
	if flag {
		students, err = service.GetStudentsFromMongoDB()
		if err != nil {
			return nil, fmt.Errorf("error fetching students from MongoDB: %v", err)
		}
	} else {
		students, err = service.GetStudentsFromPostgreSQL()
		if err != nil {
			return nil, fmt.Errorf("error fetching students from PostgreSQL: %v", err)
		}
	}
	studentResponses := make([]Response.StudentResponse, 0, len(students))
	for _, student := range students {
		studentResponses = append(studentResponses, Response.FromModel(student))
	}
	return studentResponses, nil
}

// GetStudentByID fetches a Student by ID from the selected data source
func (m *StudentManager) GetStudentByID(flag bool, id string) ([]Response.StudentResponse, error) {
	var studentResponses []Response.StudentResponse
	if flag {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf("invalid ObjectID format: %v", err)
		}
		student, err := service.GetStudentByIDFromMongoDB(objectID)
		if err != nil {
			return nil, fmt.Errorf("error fetching student from MongoDB: %v", err)
		}
		studentResponses = append(studentResponses, Response.FromModel(student))
	} else {
		student, err := service.GetStudentByIDFromPostgreSQL(id)
		if err != nil {
			return nil, fmt.Errorf("error fetching student from PostgreSQL: %v", err)
		}
		studentResponses = append(studentResponses, Response.FromModel(student))
	}
	return studentResponses, nil
}

// CreateStudent creates a new Student in the selected data source
func (m *StudentManager) CreateStudent(flag bool, student Request.StudentRequest) ([]Response.StudentResponse, error) {
	students := student.ToModel()
	var studentResponses []Response.StudentResponse
	if flag {
		insertedStudent, err := service.InsertStudent(students)
		if err != nil {
			return nil, fmt.Errorf("error inserting student into MongoDB: %v", err)
		}
		studentResponses = append(studentResponses, Response.FromModel(insertedStudent))
	} else {
		insertedStudent, err := service.InsertStudentps(students)
		if err != nil {
			return nil, fmt.Errorf("error inserting student into MongoDB :%v", err)
		}
		studentResponses = append(studentResponses, Response.FromModel(insertedStudent))
	}
	return studentResponses, nil
}

// UpdateStudent updates an existing Student in the selected data source
func (m *StudentManager) UpdateStudent(flag bool, id string, student Request.StudentRequest) ([]Response.StudentResponse, error) {
	students := student.ToModel()
	var studentResponses []Response.StudentResponse
	if flag {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf("invalid ObjectID format: %v", err)
		}
		updatedStudent, err := service.UpdateStudentInMongoDB(objectID, students)
		if err != nil {
			return nil, fmt.Errorf("invalid ObjectID format: %v", err)
		}
		studentResponses = append(studentResponses, Response.FromModel(updatedStudent))
	} else {
		updatedStudent, err := service.UpdateStudentInPostgreSQL(id, students)
		if err != nil {
			return nil, fmt.Errorf("error updating student in PostgreSQL: %v", err)
		}
		studentResponses = append(studentResponses, Response.FromModel(updatedStudent))
	}
	return studentResponses, nil
}

// DeleteStudent deletes a Student from the selected data source
func (m *StudentManager) DeleteStudent(flag bool, id string) error {
	if flag {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		return service.DeleteStudent(objectID)
	} else {
		return service.DeleteStudents(id)
	}
}
