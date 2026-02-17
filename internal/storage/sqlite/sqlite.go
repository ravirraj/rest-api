package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ravirraj/rest-api/internal/config"
	types "github.com/ravirraj/rest-api/internal/type"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {

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

	stmt, err := s.Db.Prepare("INSERT INTO students (name,email,age) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err

	}

	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {

		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("Query error : %w", err)
	}

	return student, nil
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id,name,email,age FROM students")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
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

		students = append(students, student)
	}
	return students, nil
}


func(s *Sqlite) DeleteStudentById(id int64)(int64,error) {
	stmt , err := s.Db.Prepare(`DELETE FROM students WHERE id = ?`)
	if err!= nil {
		return  0,err
	}
	defer stmt.Close()

	result , err := stmt.Exec(id)
	if err!= nil {
		return  0,err
	}

	return result.RowsAffected()
}

func (s *Sqlite ) UpdateStudentInfo(id int64 , input types.UpdateStudent) error {
	query := "UPDATE students SET "
	args := []interface{}{}
	field := []string{}

	if input.Name != nil {
		field = append(field, "name = ?")
		args = append(args, *input.Name)
	}
	if input.Email != nil {
		field = append(field, "email = ?")
		args = append(args, *input.Email)
	}
	if input.Age != nil {
		field = append(field, "age = ?")
		args = append(args, *input.Age)
	}
	if len(field) == 0 {
		return  fmt.Errorf("NO FIELD TO UPDATE \n")
	}

	query+= strings.Join(field, ", ")
	query+= "WHERE id = ?"

	args = append(args, id)

	_,err := s.Db.Exec(query,args...)
	fmt.Println(err)
	return  err
}