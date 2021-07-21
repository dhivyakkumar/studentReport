package main

import (
	"fmt"
	"net/http"
	"studentReports/src/controller"
	"studentReports/src/driver"
	"studentReports/src/repo"
)

func main() {

	fmt.Println("Starting the server")
	db := driver.OpenDBConnection()
	studentRepo := repo.NewStudentRepo(db)
	c := controller.NewController(studentRepo)
	r := controller.Router(c)

	defer db.Close()

	http.ListenAndServe(":8081", r)
}
