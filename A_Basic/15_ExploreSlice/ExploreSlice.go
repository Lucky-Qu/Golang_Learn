package main

import "fmt"

// •编写一个程序：
// 。通过循环，持续的将元素追加到 slice 里
// •在 slice 容量发生变化的时候打印出它的容量
// •请判断 append函数在底层数组的空间被填满之后，是否会将数组的容量增加一倍？
func main() {
	var s = make([]string, 0)
	capacity := cap(s)
	for i := 0; i < 2000; i++ {
		s = append(s, "")
		if capacity != cap(s) {
			fmt.Printf("发生容量改变,原容量为%d,现容量为%d\n", capacity, cap(s))
			capacity = cap(s)
		}
	}
}
