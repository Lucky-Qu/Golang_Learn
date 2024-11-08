package main

import (
	"fmt"
	"strings"
)

// 编写一个函数：
// •它可以统计文本字符串中不同单词的出现频率，并返回一个词频 map。
// • 该函数需要将文本转换为小写字母并移除包含的标点符号
// • strings 包中的 Fields、ToLower、Trim 函数应该对此有所帮助
// •使用该函数统计下文中各单词出现的频率，打印出现不止一次的单词和词频：
func count(str string) (result map[string]int) {
	str = strings.Trim(str, ",.!?—'-——")
	str = strings.ToLower(str)
	result = make(map[string]int, 10)
	for _, j := range strings.Fields(str) {
		result[j]++
	}
	return result
}
func post(m map[string]int) {
	for k, v := range m {
		if v > 1 {
			fmt.Println(k, ":", v)
		}
	}
}
func main() {
	var str = "The Republican nominee suggested on Thursday that divine intervention in next week’s election would reveal him as the rightful winner of even Democratic bastions like California.\n\nOn one level, Trump’s comment shows how his false election fraud claims have entered the outer realms of absurdity.\n\nBut this goes beyond hyperbole. Trump who altered reality for tens of millions of Americans by claiming he was cheated out of power four years ago  is creating a sinister threat to the 2024 election and spinning a legacy of fractured trust that could taint presidential votes long after he’s left the stage. The election fraud claims that Trump most notably supercharged in 2020 to salve his humiliation at losing to Joe Biden are already at high intensity this year."
	result := count(str)
	post(result)
}
