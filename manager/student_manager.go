package manager

import (
	models "student-teacher-api/model"
	"student-teacher-api/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StudentManager handles business logic for Students
type StudentManager struct{}

// GetStudents fetches all Students based on the data source
func (m *StudentManager) GetStudents(flag bool) ([]models.Student, error) {
	if flag {
		return service.GetStudentsFromMongoDB()
	} else {
		return service.GetStudentsFromPostgreSQL()
	}
}

// GetStudentByID fetches a Student by ID from the selected data source
func (m *StudentManager) GetStudentByID(flag bool, id string) (models.Student, error) {
	if flag {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return models.Student{}, err
		}
		return service.GetStudentByIDFromMongoDB(objectID)
	} else {
		return service.GetStudentByIDFromPostgreSQL(id)
	}
}

// CreateStudent creates a new Student in the selected data source
func (m *StudentManager) CreateStudent(flag bool, student models.Student) (models.Student, error) {
	if flag {
		return service.InsertStudent(student)
	} else {
		return service.InsertStudentps(student)
	}
}

// UpdateStudent updates an existing Student in the selected data source
func (m *StudentManager) UpdateStudent(flag bool, id string, student models.Student) (models.Student, error) {
	if flag {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return models.Student{}, err
		}
		updatedStudent, err := service.UpdateStudentInMongoDB(objectID, student)
		if err != nil {
			return models.Student{}, err
		}
		return updatedStudent, nil
	} else {
		updatedStudent, err := service.UpdateStudentInPostgreSQL(id, student)
		if err != nil {
			return models.Student{}, err
		}
		return updatedStudent, nil
	}
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
