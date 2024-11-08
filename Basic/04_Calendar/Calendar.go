package main

import (
	"fmt"
	"math/rand"
	"os"
)

// 实现随机生成一个年份，然后根据生成的年份自动匹配每月的天数，最后输出年，月，日，并最后随机生成十个数值并输出
func main() {
	for i := 0; i < 10; i++ {
		year := rand.Intn(3000) + 1
		month := rand.Intn(12) + 1
		var day int
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			day = rand.Intn(31)
		case 4, 6, 9, 11:
			day = rand.Intn(30)
		case 2:
			if year%4 != 0 {
				day = rand.Intn(28)
			} else {
				day = rand.Intn(29)
			}
		default:
			os.Exit(1)
		}
		fmt.Printf("年:%d月:%d日:%d\n", year, month, day)
	}
}
