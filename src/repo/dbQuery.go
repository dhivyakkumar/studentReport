package repo

import (
	"database/sql"
	"fmt"
	"sort"
	"studentReports/src/model"
)

type StudentRepo struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo{
	return &StudentRepo{db: db}
}

func (studRepo StudentRepo)GetStudents() []model.Student {
	var student model.Student
	var studentList []model.Student

	rows, err := studRepo.db.Query("SELECT * from studs")
	if err != nil {
		fmt.Errorf("Failed to select statement %v", err)
	}

	for rows.Next() {
		rows.Scan(&student.ID, &student.Name, &student.Subject1, &student.Subject2, &student.Total, &student.Avg, &student.Rank)
		studentList = append(studentList, student)
	}

	return studentList
}

func (studRepo StudentRepo)GetStudentInfo(id int) model.Student {
	var student model.Student

	row := studRepo.db.QueryRow("SELECT * FROM studs WHERE id=$1", id)

	err := row.Scan(&student.ID, &student.Name, &student.Subject1, &student.Subject2, &student.Total, &student.Avg, &student.Rank)
	if err != nil {
		fmt.Errorf("Failed to select statement %v", err)
	}

	return student
}

func (studRepo StudentRepo)CreateStudentInfo(student model.Student) model.Student{
	studRepo.db.QueryRow("INSERT INTO studs(name, subject1, subject2) VALUES($1, $2, $3) RETURNING id",
		student.Name, student.Subject1, student.Subject2).Scan(&student.ID)

	return student
}

func (studRepo StudentRepo)UpdateStudentInfo(id int, stud model.Student) model.Student{

	studRepo.db.QueryRow("UPDATE studs SET name=$1, subject1=$2, subject2=$3 WHERE id=$4", stud.Name, stud.Subject1, stud.Subject2, id)

	stud.ID = id
	return stud
}

func (studRepo StudentRepo)RemoveStudentInfo(id int){

	studRepo.db.QueryRow("DELETE FROM studs WHERE id=$1", id)
}

func (studRepo StudentRepo)GetResultInfo() []model.Student{
	var studList []model.Student
	var stud model.Student


	rows, err:= studRepo.db.Query("SELECT * FROM studs")
	if err!=nil{
		fmt.Errorf("Failed to execute select %v", err)
	}

	for rows.Next(){
		rows.Scan(&stud.ID, &stud.Name, &stud.Subject1, &stud.Subject2, &stud.Total, &stud.Avg, &stud.Rank)
		stud.Total = stud.Subject1 + stud.Subject2
		stud.Avg = float32(stud.Total / 2)
		studList = append(studList, stud)
	}

	studList= sortList(studList)
	for i:=0;i< len(studList);i++{
		studList[i].Rank = i+1
	}

	studRepo.insertStudentInfo(studList)
	return studList
}

func sortList(studList []model.Student) []model.Student{
	sort.SliceStable(studList, func(i, j int) bool {
		return studList[i].Avg > studList[j].Avg
	})
	return studList
}

func (studRepo StudentRepo)insertStudentInfo(studList []model.Student){

	for _,stud:=range studList{
		studRepo.db.Query("UPDATE studs SET total=$1, avg=$2, rank=$3 WHERE id=$4",stud.Total, stud.Avg, stud.Rank,stud.ID)
	}
}