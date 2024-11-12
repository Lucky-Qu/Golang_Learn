package main

import "fmt"

// •亚瑟被一位骑士挡住了去路。正如 leftHand *item 变量的值为 nil 所示，这位英雄手上正空无一物。
// • 请实现一个拥有 pickup（i *item） 和 give（to *character）等方法的character 结构，然后使用你在本节学到的知识编写一个脚本，使得
// 亚瑟可以拿起一件物品并将其交给骑士，与此同时每个动作打印出适当描述。
type item struct {
	name string
}
type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c.leftHand != nil {
		fmt.Printf("%s手上已经有东西了\n", c.name)
	}
	c.leftHand = i
	fmt.Printf("%s拿起了%s\n", c.name, i.name)
}
func (c *character) give(to *character) {
	if c.leftHand == nil {
		fmt.Printf("%s手上是空的！\n", c.name)
	}
	if to.leftHand != nil {
		fmt.Printf("%s手上有东西！\n", to.name)
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%s把%s递给了%s\n", c.name, to.leftHand.name, to.name)
}
func main() {
	arthur := character{name: "arthur", leftHand: nil}
	knight := character{name: "knight", leftHand: nil}
	sword := item{name: "sword"}
	arthur.pickup(&sword)
	arthur.give(&knight)
}
