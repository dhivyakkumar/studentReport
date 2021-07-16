package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"studentReports/src/model"
	"studentReports/src/repo"
)

func getAllStudentList(w http.ResponseWriter, r *http.Request){

  	var studentList []model.Student
	studentList = repo.GetStudents()
	json.NewEncoder(w).Encode(studentList)
	w.WriteHeader(http.StatusOK)
}

func getStudent(w http.ResponseWriter, r *http.Request){
	param:=mux.Vars(r)
	id,err := strconv.Atoi(param["id"])
	if err!=nil{
		fmt.Errorf("Failed to convert type", err)
	}

	var student model.Student

	student = repo.GetStudentInfo(id)

	if student.ID!=0	{
	json.NewEncoder(w).Encode(student)
	w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func createStudent(w http.ResponseWriter, r *http.Request){

	var student model.Student
	var newStudent model.Student
	json.NewDecoder(r.Body).Decode(&student)

	newStudent = repo.CreateStudentInfo(student)
	json.NewEncoder(w).Encode(newStudent)
	w.WriteHeader(http.StatusCreated)
}

func updateStudent(w http.ResponseWriter, r *http.Request){

	var stud model.Student
	var studUpdated model.Student
	param:=mux.Vars(r)
	id, err:=strconv.Atoi(param["id"])
	if err!=nil{
		fmt.Errorf("Failed to convert type %v", err)
	}

	json.NewDecoder(r.Body).Decode(&stud)

	studUpdated = repo.UpdateStudentInfo(id, stud)

	json.NewEncoder(w).Encode(studUpdated)
	w.WriteHeader(http.StatusOK)
}

func removeStudent(w http.ResponseWriter, r *http.Request){

	param:=mux.Vars(r)
	id,err:=strconv.Atoi(param["id"])
	if err!=nil{
		fmt.Errorf("Failed to convert %v", err)
	}

	repo.RemoveStudentInfo(id)
	w.WriteHeader(http.StatusOK)
}

func getResult(w http.ResponseWriter, r *http.Request) {
	var studList []model.Student
	studList = repo.GetResultInfo()

	json.NewEncoder(w).Encode(studList)
	w.WriteHeader(http.StatusOK)
}