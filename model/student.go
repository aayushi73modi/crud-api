package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID           string `bson:"_id,omitempty" json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid()" `
	Student_name string `bson:"student_name" json:"student_name" gorm:"column:student_name" validate:"required"`
	Age          int    `bson:"age" json:"age" gorm:"column:age" validate:"required"`
	Class        string `bson:"class" json:"class" gorm:"column:class" validate:"required"`
}

func (s *Student) SetMongoID() {
	if s.ID == "" {
		s.ID = primitive.NewObjectID().Hex()
	}
}

func (s *Student) GenerateUUID() {
	if s.ID == "" {
		s.ID = uuid.New().String() // For PostgreSQL
	}
}
