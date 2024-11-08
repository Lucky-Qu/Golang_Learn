package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 康威生命游戏
const (
	width  = 80
	height = 15
)

type Universe [][]bool

// NewUniverse 编写NewUniverse函数以用make生成一个长宽为width和height的Universe
func NewUniverse(width, height int) (universe Universe) {
	universe = make(Universe, height)
	for i := range universe {
		universe[i] = make([]bool, width)
	}
	return universe
}

// Seed 编写Seed方法用于随机激活宇宙中25%的细胞
func (u Universe) Seed() {
	for _, v := range u {
		for i := range v {
			if rand.Intn(100) < 25 {
				v[i] = true
			}
		}
	}
}

// Show 编写show方法为universe利用fmt包打印目前状态死了的用星号表示，存活的用空格表示
func (u Universe) Show() {
	for i := range u {
		for _, k := range u[i] {
			if k {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Alive 判断指定位置细胞是否存活并实现回绕
func (u Universe) Alive(x, y int) bool {
	x = (x + height) % height
	y = (y + width) % width
	return u[x][y]
}

// Neighbors 统计临近细胞存活数
func (u Universe) Neighbors(x, y int) int {
	var result int
	if u.Alive(x, y) {
		result = -1
	} else {
		result = 0
	}
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if u.Alive(x+i, y+j) {
				result++
			}
		}
	}
	return result
}

// Next 实现规则
// 当一个存活细胞邻近的存活细胞少于2个时，该细胞死亡。
// 当一个存活细胞邻近有2个或3个存活细胞时，该细胞将延续至下一世代。
// 当一个存活细胞邻近有多于3个存活细胞时，该细胞死亡。
// 当一个死亡细胞邻近正好有3个存活细胞时，该细胞存活。
func (u Universe) Next(x, y int) bool {
	switch u.Neighbors(x, y) {
	case 0, 1:
		return false
	case 2:
		return u.Alive(x, y)
	case 3:
		if !u.Alive(x, y) {
			return true
		}
		return u.Alive(x, y)
	case 4, 5, 6, 7, 8:
		return false
	}
	fmt.Println("出错了！", u.Neighbors(x, y))
	return false
}

// Step 创建平行世界进行操作
func Step(a, b Universe) {
	for i := range a {
		for j := range a[i] {
			b[i][j] = a.Next(i, j)
		}
	}
}
func main() {
	a := NewUniverse(width, height)
	b := a
	a.Seed()
	for true {
		a.Show()
		Step(a, b)
		a, b = b, a
		time.Sleep(time.Second * 1)
		fmt.Println("================================================================================")
	}
}
