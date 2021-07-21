package controller

import "github.com/gorilla/mux"

func Router(c *Controller) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/students", c.getAllStudentList).Methods("GET")
	r.HandleFunc("/api/student/{id}", c.getStudent).Methods("GET")
	r.HandleFunc("/api/student", c.createStudent).Methods("POST")
	r.HandleFunc("/api/student/{id}", c.updateStudent).Methods("PUT")
	r.HandleFunc("/api/student/{id}", c.removeStudent).Methods("DELETE")
	r.HandleFunc("/api/students/result", c.getResult).Methods("GET")

	return r
}
