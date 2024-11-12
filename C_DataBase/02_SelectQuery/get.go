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

func getMany(id int) ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT id, name, age, sex, hobbies FROM users WHERE id>?", id)
	for rows.Next() {
		u := User{}
		rows.Scan(&u.Id, &u.Name, &u.Age, &u.Sex, &u.Hobbies)
		users = append(users, u)
	}
	return users, err
}
