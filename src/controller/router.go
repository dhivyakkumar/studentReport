package controller

import "github.com/gorilla/mux"

func Router() *mux.Router{
	r:=mux.NewRouter()

	r.HandleFunc("/api/students",getAllStudentList).Methods("GET")
	r.HandleFunc("/api/student/{id}",getStudent).Methods("GET")
	r.HandleFunc("/api/student",createStudent).Methods("POST")
	r.HandleFunc("/api/student/{id}",updateStudent).Methods("PUT")
	r.HandleFunc("/api/student/{id}",removeStudent).Methods("DELETE")
	r.HandleFunc("/api/students/result",getResult).Methods("GET")

	return r
}
