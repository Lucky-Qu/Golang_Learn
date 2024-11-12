package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:secret@/mysql")
	if err != nil {
		fmt.Println(err)
	}
	err = db.PingContext(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully connected to the database")
}
