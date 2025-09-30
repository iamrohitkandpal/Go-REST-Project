package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/iamrohitkandpal/Go-REST-Project/internal/types"

	"github.com/iamrohitkandpal/Go-REST-Project/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	statement, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer statement.Close()
	res, err := statement.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}


func (s *Sqlite) GetStudentById(id int64) (types.Student, error)  {
	statement, err := s.Db.Prepare("SELECT id, name, email, age FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	} 

	defer statement.Close()

	var student types.Student

	err = statement.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}

	return student, nil
}

func (s *Sqlite) GetAllStudents() ([]types.Student, error) {
	statement, err := s.Db.Prepare("SELECT id, name, email, age FROM students LIMIT 5")

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)

		if err != nil {
			return nil, err
		}

		students =append(students, student)
	}

	return students, nil
}


func (s *Sqlite) UpdateStudentById(id int64, name string, email string, age int) (types.Student, error) {
    // First check if student exists
    _, err := s.GetStudentById(id)
    if err != nil {
        return types.Student{}, err
    }

    // Update the student
    statement, err := s.Db.Prepare("UPDATE students SET name = ?, email = ?, age = ? WHERE id = ?")
    if err != nil {
        return types.Student{}, err
    }
    defer statement.Close()

    _, err = statement.Exec(name, email, age, id)
    if err != nil {
        return types.Student{}, err
    }

    // Return updated student
    updatedStudent := types.Student{
        Id:    id,
        Name:  name,
        Email: email,
        Age:   age,
    }

    return updatedStudent, nil
}


func (s *Sqlite) DeleteStudentById(id int64) (types.Student, error) {
	student, err := s.GetStudentById(id)
	if err != nil {
		return types.Student{}, err
	}

	statement, err := s.Db.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		return types.Student{}, err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return types.Student{}, err
	}

	return student, nil
}