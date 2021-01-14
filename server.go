package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func main() {

	db, err := sql.Open("godror", `user="akshay" password="password123" connectString="localhost:1521/ORCL"
                               libDir="/Users/akshaybavalekar/Downloads/instantclient_19_8"`)
	if err != nil {
		fmt.Println(err)
		return
	}

	dbQuery := "CREATE TABLE issues(ID VARCHAR2(20), ISSUE VARCHAR2(20))"
	rows, err := db.Query(dbQuery)
	if err != nil {
		fmt.Println(".....Error processing query")
		fmt.Println(err)
		return
	}
	defer rows.Close()
	fmt.Println("... Connected ")

}
