package repo

import (
	"database/sql"
	"fmt"
	"studentReports/src/model"
)
type Calculate interface {
	CalculateRank(studList []model.Student) []model.Student
}

type StudentRepo struct {
	db *sql.DB
	cal Calculate
}

func NewStudentRepo(db *sql.DB, cal Calculate) *StudentRepo {
	return &StudentRepo{
		db: db,
		cal: cal,
	}
}

func (studRepo StudentRepo) GetStudents() ([]model.Student, error) {
	var student model.Student
	var studentList []model.Student

	rows, err := studRepo.db.Query("SELECT * from studs")
	if err != nil {
		return []model.Student{}, fmt.Errorf("Failed to select statement %v", err)
	}

	for rows.Next() {
		rows.Scan(&student.ID, &student.Name, &student.Subject1, &student.Subject2, &student.Total, &student.Avg, &student.Rank)
		studentList = append(studentList, student)
	}

	return studentList, nil
}

func (studRepo StudentRepo) GetStudentInfo(id int) (model.Student, error) {
	var student model.Student

	row := studRepo.db.QueryRow("SELECT * FROM studs WHERE id=$1", id)
	err := row.Scan(&student.ID, &student.Name, &student.Subject1, &student.Subject2, &student.Total, &student.Avg, &student.Rank)
	if err != nil {
		return model.Student{}, fmt.Errorf("Failed to select statement %v", err)
	}

	return student, nil
}

func (studRepo StudentRepo) CreateStudentInfo(student model.Student) (model.Student, error) {
	err := studRepo.db.QueryRow("INSERT INTO studs(name, subject1, subject2) VALUES($1, $2, $3) RETURNING id",
		student.Name, student.Subject1, student.Subject2).Scan(&student.ID)
	if err != nil {
		return model.Student{}, fmt.Errorf("Failed to create statement %v", err)
	}

	return student, nil
}

func (studRepo StudentRepo) UpdateStudentInfo(id int, stud model.Student) (model.Student, error) {
	err := studRepo.db.QueryRow("UPDATE studs SET name=$1, subject1=$2, subject2=$3 WHERE id=$4", stud.Name, stud.Subject1, stud.Subject2, id)
	if err != nil {
		return model.Student{}, fmt.Errorf("Failed to update statement %v", err)
	}
	stud.ID = id

	return stud, nil
}

func (studRepo StudentRepo) RemoveStudentInfo(id int) error {
	err := studRepo.db.QueryRow("DELETE FROM studs WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("Failed to execute delete query %v", err)
	}

	return nil
}

func (studRepo StudentRepo) GetResultInfo() ([]model.Student, error) {
	var studList []model.Student
	var stud model.Student

	rows, err := studRepo.db.Query("SELECT * FROM studs")
	if err != nil {
		return []model.Student{}, fmt.Errorf("Failed to execute select %v", err)
	}

	for rows.Next() {
		rows.Scan(&stud.ID, &stud.Name, &stud.Subject1, &stud.Subject2, &stud.Total, &stud.Avg, &stud.Rank)
		studList = append(studList, stud)
	}

	studData := studRepo.cal.CalculateRank(studList)

	err = studRepo.updateStudentResults(studData)
	if err != nil {
		return []model.Student{}, fmt.Errorf("Failed to update student result %v", err)
	}

	return studData, nil
}

func (studRepo StudentRepo) updateStudentResults(studList []model.Student) error {
	for _, stud := range studList {
		_, err := studRepo.db.Query("UPDATE studs SET total=$1, avg=$2, rank=$3 WHERE id=$4", stud.Total, stud.Avg, stud.Rank, stud.ID)
		if err != nil {
			return fmt.Errorf("Failed to update student result %v", err)
		}
	}

	return nil
}
