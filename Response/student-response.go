package Response

import (
	models "student-teacher-api/model"
)

// studentResponse is the response structure for a student
type StudentResponse struct {
	ID           string `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid()"`
	Student_name string `json:"student_name"`
	Age          int    `json:"age"`
	Class        string `json:"class"`
}

// FromModel converts a Student model to a studentResponse
func FromModel(student models.Student) StudentResponse {
	return StudentResponse{
		ID:           student.ID,
		Student_name: student.Student_name,
		Age:          student.Age,
		Class:        student.Class,
	}
}
