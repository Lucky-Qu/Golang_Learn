package main

import "fmt"

// •写一个程序，把字符串转化为布尔类型：
// • "true", "yes", "1" -> true
// • "false", "no", "0"  -> false
// • 针对其它值，显示错误信息
func main() {
	var str = "1"
	var x = []string{"true", "yes", "1"}
	var y = []string{"false", "no", "0"}
	var result bool
	for _, r := range x {
		if str == r {
			result = true
			fmt.Println(result)
			break
		}
	}
	for _, r := range y {
		if str == r {
			result = false
			fmt.Println(result)
			break
		}
	}
	//应该整个函数，上面有结果就return了，懒得改了
	fmt.Println("错误信息")
}
