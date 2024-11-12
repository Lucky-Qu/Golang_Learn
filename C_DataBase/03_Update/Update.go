package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// User 定义用户结构体
type User struct {
	Id   int
	Name string
	Age  int
	Sex  string
}

var db *sql.DB

// getOne 执行SELECT操作，返回查询的值
func getOne(id int) (u User, err error) {
	err = db.QueryRow("SELECT Id, Name, Age, Sex FROM updatetest WHERE id=?", id).Scan(&u.Id, &u.Name, &u.Age, &u.Sex)
	return u, err
}
func update(u User) (err error) {
	_, err = db.Exec("UPDATE updatetest SET Name=?,Age=?,Sex=? WHERE Id=?", "xiaoshi", "10", "female", u.Id)
	return err
}
func main() {
	var (
		err error
		u   User
	)
	db, err = sql.Open("mysql", "root:Secret@/Learn")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = db.PingContext(context.Background())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("连接数据库成功！")
	fmt.Println("下一步执行查找操作")
	u, err = getOne(1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("执行查询操作成功！查找到的值为:", u)
	fmt.Println("下一步执行修改操作")
	err = update(u)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("修改操作执行成功！")
	fmt.Println("下一步执行查找操作！")
	u, err = getOne(1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("执行查找操作成功！查找到的值为：", u)
}
