package repo

import (
	"fmt"
	"sort"
	"studentReports/src/driver"
	"studentReports/src/model"
)

func GetStudents() []model.Student {
	var student model.Student
	var studentList []model.Student

	rows, err := driver.DB.Query("SELECT * from studs")
	if err != nil {
		fmt.Errorf("Failed to select statement %v", err)
	}

	for rows.Next() {
		rows.Scan(&student.ID, &student.Name, &student.Subject1, &student.Subject2, &student.Total, &student.Avg, &student.Rank)
		studentList = append(studentList, student)
	}

	return studentList
}

func GetStudentInfo(id int) model.Student {
	var student model.Student

	row := driver.DB.QueryRow("SELECT * FROM studs WHERE id=$1", id)

	err := row.Scan(&student.ID, &student.Name, &student.Subject1, &student.Subject2, &student.Total, &student.Avg, &student.Rank)
	if err != nil {
		fmt.Errorf("Failed to select statement %v", err)
	}

	return student
}

func CreateStudentInfo(student model.Student) model.Student{
	driver.DB.QueryRow("INSERT INTO studs(name, subject1, subject2) VALUES($1, $2, $3) RETURNING id",
		student.Name, student.Subject1, student.Subject2).Scan(&student.ID)

	return student
}

func UpdateStudentInfo(id int, stud model.Student) model.Student{

	driver.DB.QueryRow("UPDATE studs SET name=$1, subject1=$2, subject2=$3 WHERE id=$4", stud.Name, stud.Subject1, stud.Subject2, id)

	stud.ID = id
	return stud
}

func RemoveStudentInfo(id int){

	driver.DB.QueryRow("DELETE FROM studs WHERE id=$1", id)
}

func GetResultInfo() []model.Student{
	var studList []model.Student
	//var sortedStudList []model.Student
	var stud model.Student


	rows, err:= driver.DB.Query("SELECT * FROM studs")
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

	insertStudentInfo(studList)
	return studList
}

func sortList(studList []model.Student) []model.Student{
	sort.SliceStable(studList, func(i, j int) bool {
		return studList[i].Avg > studList[j].Avg
	})
	return studList
}

func insertStudentInfo(studList []model.Student){

	for _,stud:=range studList{
		driver.DB.Query("UPDATE studs SET total=$1, avg=$2, rank=$3 WHERE id=$4",stud.Total, stud.Avg, stud.Rank,stud.ID)
	}
}