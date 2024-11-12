package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func getOne(id int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, age, sex, hobbies FROM users WHERE id=?", id).Scan(&user.Id, &user.Name, &user.Age, &user.Sex, &user.Hobbies)
	if err == sql.ErrNoRows {
		return user, fmt.Errorf("user with ID %d not found", id)
	}
	if err != nil {
		return user, err
	}
	return user, nil
}
