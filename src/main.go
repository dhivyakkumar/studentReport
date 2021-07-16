package main

import (
	"fmt"
	"net/http"
	"studentReports/src/controller"
	"studentReports/src/driver"
)



func main(){

	fmt.Println("Starting the server")


	r:=controller.Router()
	db:= driver.OpenDBConnection()
	defer db.Close()
	http.ListenAndServe(":8081",r)
}
