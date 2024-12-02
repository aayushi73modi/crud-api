package Request

import (
	models "student-teacher-api/model"
)

// StudentRequest is the request structure for creating and updating a Student
type StudentRequest struct {
	Student_name string `bson:"student_name" json:"student_name" validate:"required" gorm:"student_name"`
	Age          int    `bson:"age" json:"age" validate:"required" gorm:"age"`
	Class        string `bson:"class" json:"class" validate:"required" gorm:"class"`
}

// ToModel converts a StudentRequest to a Student model
func (req *StudentRequest) ToModel() models.Student {
	return models.Student{
		Student_name: req.Student_name,
		Age:          req.Age,
		Class:        req.Class,
	}
}
