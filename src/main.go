package main

import (
	"fmt"
	"net/http"
	"studentReports/src/controller"
	"studentReports/src/driver"
	"studentReports/src/repo"
	"studentReports/src/resultCalculator"
)

func main() {

	fmt.Println("Starting the server")
	db := driver.OpenDBConnection()
	cal:=resultCalculator.NewCalculator()
	studentRepo := repo.NewStudentRepo(db, cal)
	c := controller.NewController(studentRepo)
	r := controller.Router(c)

	defer db.Close()

	http.ListenAndServe(":8081", r)
}
