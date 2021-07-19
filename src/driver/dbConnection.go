package driver

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
	"os"
)

func init()  {
	gotenv.Load()
}

func OpenDBConnection() *sql.DB{
	fmt.Println("Connecting to DB!!!")

	pgURL,err :=pq.ParseURL(os.Getenv("SQL"))
	if err!=nil{
		fmt.Errorf("Failed to parse file", err)
	}

	db, _ := sql.Open("postgres", pgURL)

	err= db.Ping()
	if err!=nil{
		fmt.Errorf("Failed to ping db", err)
	}

	return db
}