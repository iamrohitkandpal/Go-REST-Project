package storage

import "github.com/iamrohitkandpal/Go-REST-Project/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetAllStudents() ([]types.Student, error)
	UpdateStudentById(id int64) (types.Student, error)
	DeleteStudentById(id int64) (types.Student, error)
}