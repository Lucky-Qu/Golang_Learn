package main

import (
	"fmt"
	"math/rand"
)

// 写一个程序，把布尔类型转化为整数类型
func main() {
	r := []bool{true, false}
	x1 := r[rand.Intn(len(r))]
	var x2 int
	if x1 {
		x2 = 1
	} else {
		x2 = 0
	}
	fmt.Println(x1, x2)
}
