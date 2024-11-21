package manager

import (
	"errors"
	models "student-teacher-api/model"
	"student-teacher-api/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StudentManager handles business logic for Students
type StudentManager struct{}

// GetStudents fetches all Students based on the data source
func (m *StudentManager) GetStudents(flag bool) ([]models.Student, error) {
	switch flag {
	case true: //1
		return service.GetStudentsFromMongoDB()
	case false: //0
		return service.GetStudentsFromPostgreSQL()
	default:
		return nil, errors.New("invalid flag parameter")
	}
}

// GetStudentByID fetches a Student by ID from the selected data source
func (m *StudentManager) GetStudentByID(flag bool, id string) (models.Student, error) {
	switch flag {
	case true:
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return models.Student{}, err
		}
		return service.GetStudentByIDFromMongoDB(objectID)
	case false:
		return service.GetStudentByIDFromPostgreSQL(id)
	default:
		return models.Student{}, errors.New("invalid flag parameter")
	}
}

// CreateStudent creates a new Student in the selected data source
func (m *StudentManager) CreateStudent(flag bool, student models.Student) (models.Student, error) {
	switch flag {
	case true:
		return service.InsertStudent(student)
	case false:
		return service.InsertStudentps(student)
	default:
		return models.Student{}, errors.New("invalid flag parameter")
	}
}

// UpdateStudent updates an existing Student in the selected data source
func (m *StudentManager) UpdateStudent(flag bool, id string, student models.Student) (models.Student, error) {
	switch flag {
	case true: // MongoDB
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return models.Student{}, err
		}
		updatedStudent, err := service.UpdateStudentInMongoDB(objectID, student)
		if err != nil {
			return models.Student{}, err
		}
		return updatedStudent, nil
	case false: // PostgresSQL
		updatedStudent, err := service.UpdateStudentInPostgreSQL(id, student)
		if err != nil {
			return models.Student{}, err
		}
		return updatedStudent, nil
	default:
		return models.Student{}, errors.New("invalid flag parameter")
	}
}

// DeleteStudent deletes a Student from the selected data source
func (m *StudentManager) DeleteStudent(flag bool, id string) error {
	switch flag {
	case true:
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		return service.DeleteStudent(objectID)
	case false:
		return service.DeleteStudents(id)
	default:
		return errors.New("invalid flag parameter")
	}
}
