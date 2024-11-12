// 1.编写一个流水线部件（一个 goroutine），他需要记住前面出现的所有值，并且只有在值之前从未出现过的情况下才会将值传递至流水线的下一个阶段。假定第一个值永远不是空字符串。
// 2. 编写一个流水线部件，它接收字符串并将它们拆分成单词，然后向流水线的下一阶段一个接一个的发送这些单词。
// •可以使用 strings.Fields 函数
package main

import (
	"fmt"
	"strings"
)

func Gopher1(c1 chan string, c2 chan string) {
	str := <-c1
	close(c1)
	result := strings.Fields(str)
	for _, v := range result {
		c2 <- v
	}
	close(c2)
}

func Gopher2(c1 chan string, c2 chan string) {
	m := make(map[string]bool)
	for v := range c1 {
		if !m[v] {
			m[v] = true
			c2 <- v
		}
	}
	close(c2)
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go Gopher1(c1, c2)
	go Gopher2(c2, c3)

	c1 <- "Hello world Hello world Hell word"

	for v := range c3 {
		fmt.Println(v)
	}
}
