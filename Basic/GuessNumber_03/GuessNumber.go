package main

import (
	"fmt"
	"math/rand"
)

// 随机一个数字，计算机猜这个数字，猜错就重猜猜到猜对为止，并附加文字提示
func main() {
	a := rand.Intn(10)
	fmt.Println("给出的数字是：", a)
	for {
		if b := rand.Intn(10); b < a {
			fmt.Println("猜小了！")
		} else if b > a {
			fmt.Println("猜大了！")
		} else {
			fmt.Println("猜对了！")
			break
		}
	}
}
