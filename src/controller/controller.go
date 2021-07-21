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
	GetStudents() ([]model.Student, error)
	GetStudentInfo(id int) (model.Student, error)
	CreateStudentInfo(student model.Student) (model.Student, error)
	UpdateStudentInfo(id int, stud model.Student) (model.Student, error)
	RemoveStudentInfo(id int) error
	GetResultInfo() ([]model.Student, error)
}

type Controller struct {
	studentRepo StudentRepoOperations
}

func NewController(sr StudentRepoOperations) *Controller {
	return &Controller{
		studentRepo: sr,
	}
}

func (ctrl Controller) getAllStudentList(w http.ResponseWriter, r *http.Request) {
	studentList, err := ctrl.studentRepo.GetStudents()
	if err != nil {
		fmt.Errorf("Failed to get students info %v", err)
	}

	json.NewEncoder(w).Encode(studentList)
	w.WriteHeader(http.StatusOK)
}

func (ctrl Controller) getStudent(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	param := mux.Vars(r)

	id, err := strconv.Atoi(param["id"])
	if err != nil {
		fmt.Errorf("Failed to convert type %v", err)
	}

	student, err = ctrl.studentRepo.GetStudentInfo(id)
	if err != nil {
		fmt.Errorf("Failed to get student info %v", err)
	}

	if student.ID != 0 {
		json.NewEncoder(w).Encode(student)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (ctrl Controller) createStudent(w http.ResponseWriter, r *http.Request) {
	var student model.Student

	json.NewDecoder(r.Body).Decode(&student)
	newStudent, err := ctrl.studentRepo.CreateStudentInfo(student)
	if err != nil {
		fmt.Errorf("Failed to create student info %v", err)
	}

	json.NewEncoder(w).Encode(newStudent)
	w.WriteHeader(http.StatusCreated)
}

func (ctrl Controller) updateStudent(w http.ResponseWriter, r *http.Request) {
	var stud model.Student
	var studUpdated model.Student
	param := mux.Vars(r)

	id, err := strconv.Atoi(param["id"])
	if err != nil {
		fmt.Errorf("Failed to convert type %v", err)
	}

	json.NewDecoder(r.Body).Decode(&stud)

	studUpdated, err = ctrl.studentRepo.UpdateStudentInfo(id, stud)
	if err != nil {
		fmt.Errorf("Failed to update student info %v", err)
	}

	json.NewEncoder(w).Encode(studUpdated)
	w.WriteHeader(http.StatusOK)
}

func (ctrl Controller) removeStudent(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	id, err := strconv.Atoi(param["id"])
	if err != nil {
		fmt.Errorf("Failed to convert %v", err)
	}

	err = ctrl.studentRepo.RemoveStudentInfo(id)
	if err != nil {
		fmt.Errorf("Failed to remove student info %v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func (ctrl Controller) getResult(w http.ResponseWriter, r *http.Request) {
	var studList []model.Student

	studList, err := ctrl.studentRepo.GetResultInfo()
	if err != nil {
		fmt.Errorf("Failed to get student result info %v", err)
	}

	json.NewEncoder(w).Encode(studList)
	w.WriteHeader(http.StatusOK)
}
