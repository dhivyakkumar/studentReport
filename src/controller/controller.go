package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"studentReports/src/model"
)

type StudentRepoOperations interface {
	GetStudents() []model.Student
	GetStudentInfo(id int) model.Student
	CreateStudentInfo(student model.Student) model.Student
	UpdateStudentInfo(id int, stud model.Student) model.Student
	RemoveStudentInfo(id int)
	GetResultInfo() []model.Student
}

type Controller struct {
	studentRepo StudentRepoOperations
}

func NewController(sr StudentRepoOperations)*Controller{
	return &Controller{
		studentRepo: sr,
	}
}

func (ctrl Controller)getAllStudentList(w http.ResponseWriter, r *http.Request){

  	var studentList []model.Student
	studentList = ctrl.studentRepo.GetStudents()
	json.NewEncoder(w).Encode(studentList)
	w.WriteHeader(http.StatusOK)
}

func (ctrl Controller)getStudent(w http.ResponseWriter, r *http.Request){
	param:=mux.Vars(r)
	id,err := strconv.Atoi(param["id"])
	if err!=nil{
		fmt.Errorf("Failed to convert type", err)
	}

	var student model.Student

	student = ctrl.studentRepo.GetStudentInfo(id)

	if student.ID!=0	{
	json.NewEncoder(w).Encode(student)
	w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (ctrl Controller)createStudent(w http.ResponseWriter, r *http.Request){

	var student model.Student
	var newStudent model.Student
	json.NewDecoder(r.Body).Decode(&student)

	newStudent = ctrl.studentRepo.CreateStudentInfo(student)
	json.NewEncoder(w).Encode(newStudent)
	w.WriteHeader(http.StatusCreated)
}

func (ctrl Controller)updateStudent(w http.ResponseWriter, r *http.Request){

	var stud model.Student
	var studUpdated model.Student
	param:=mux.Vars(r)
	id, err:=strconv.Atoi(param["id"])
	if err!=nil{
		fmt.Errorf("Failed to convert type %v", err)
	}

	json.NewDecoder(r.Body).Decode(&stud)

	studUpdated = ctrl.studentRepo.UpdateStudentInfo(id, stud)

	json.NewEncoder(w).Encode(studUpdated)
	w.WriteHeader(http.StatusOK)
}

func (ctrl Controller)removeStudent(w http.ResponseWriter, r *http.Request){

	param:=mux.Vars(r)
	id,err:=strconv.Atoi(param["id"])
	if err!=nil{
		fmt.Errorf("Failed to convert %v", err)
	}

	ctrl.studentRepo.RemoveStudentInfo(id)
	w.WriteHeader(http.StatusOK)
}

func (ctrl Controller)getResult(w http.ResponseWriter, r *http.Request) {
	var studList []model.Student
	studList = ctrl.studentRepo.GetResultInfo()

	json.NewEncoder(w).Encode(studList)
	w.WriteHeader(http.StatusOK)
}