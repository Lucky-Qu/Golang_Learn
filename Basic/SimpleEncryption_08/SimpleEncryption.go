package main

import "fmt"

// Lfdph, L vdz, Lfrqtxhuhg，每个字母向前移动3个位置，能得到什么字符串？
func main() {
	password := []rune("Lfdph, L vdz, Lfrqtxhuhg")
	for index, x := range password {
		if x == ',' || x == ' ' {
			continue
		}
		password[index] = x - 3
	}
	fmt.Println(string(password))
}
