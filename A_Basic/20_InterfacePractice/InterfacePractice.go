package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NightStartHour = 18
	NightEndHour   = 6
)

type Animal interface {
	Stringer
	Mover
	Eater
}

type Stringer interface {
	Describe() string
}

type Mover interface {
	Move() string
}

type Eater interface {
	Eat() string
}

type Monkey struct {
	Name string
}

func (m Monkey) Describe() string {
	return fmt.Sprintf("猴子：%s", m.Name)
}

func (m Monkey) Move() string {
	return fmt.Sprintf("%s正在攀爬一棵树！", m.Name)
}

func (m Monkey) Eat() string {
	return fmt.Sprintf("%s正在享用美味的香蕉！", m.Name)
}

type Dog struct {
	Name string
}

func (d Dog) Describe() string {
	return fmt.Sprintf("狗：%s", d.Name)
}

func (d Dog) Move() string {
	return fmt.Sprintf("%s正在快乐地奔跑！", d.Name)
}

func (d Dog) Eat() string {
	return fmt.Sprintf("%s正在啃食一根大骨头！", d.Name)
}

func Day(animals []Animal) {
	rand.Seed(time.Now().UnixNano())

	for hour := 0; hour < 24*3; hour++ { // 模拟三天
		if hour >= NightStartHour && hour <= NightEndHour {
			fmt.Println("夜幕降临，动物们进入梦乡...")
		} else {
			animal := animals[rand.Intn(len(animals))]
			action := rand.Intn(2)
			switch action {
			case 0:
				fmt.Println(animal.Eat())
			case 1:
				fmt.Println(animal.Move())
			}
		}
		time.Sleep(1 * time.Second) // 增加延迟，使输出更加清晰
	}
}

func main() {
	monkey := Monkey{"孙悟空"}
	dog := Dog{"旺财"}

	Day([]Animal{&monkey, &dog})
}
