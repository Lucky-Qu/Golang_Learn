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
func createUser(u User) (err error) {
	_, err = db.Exec("INSERT INTO updatetest (Id, Name, Age, Sex) VALUES (?, ?, ?, ?)", u.Id, u.Name, u.Age, u.Sex)
	return err
}
func main() {
	var (
		err error
		u   User
	)
	db, err = sql.Open("mysql", "user:password@/table")
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
	fmt.Println("查找出的数据为：", u)
	fmt.Println("下一步执行新建操作")
	u = User{
		Id:   3,
		Name: "test",
		Age:  20,
		Sex:  "male",
	}
	err = createUser(u)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("新建操作成功")
	fmt.Println("下一步执行查找操作")
	u, err = getOne(3)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("查找成功，值为：", u)
	fmt.Println("执行完毕")
}
