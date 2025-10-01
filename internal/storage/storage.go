package storage

import "github.com/iamrohitkandpal/Go-REST-Project/internal/types"

type Storage interface {
    CreateStudent(name string, email string, age int) (int64, error)
    GetStudentById(id int64) (types.Student, error)
    GetAllStudents() ([]types.Student, error)
    UpdateStudentById(id int64, name string, email string, age int) (types.Student, error)
    DeleteStudentById(id int64) (types.Student, error)
}

type AdminStorage interface {
    CreateAdmin(username string, email string, hashedPassword string) (int64, error)
    GetAdminById(id int64) (types.Admin, error)
    GetAdminByEmail(email string) (types.Admin, error)
    GetAdminByUsername(username string) (types.Admin, error)
    GetAllAdmins() ([]types.AdminResponse, error)
    UpdateAdminById(id int64, username string, email string) (types.AdminResponse, error)
    DeleteAdminById(id int64) (types.AdminResponse, error)
}