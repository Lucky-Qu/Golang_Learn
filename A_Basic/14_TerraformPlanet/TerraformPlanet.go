package main

import "fmt"

// •编写一个程序：
// •它通过给字符串 slice 中所有的行星加上“New”前缀来完成行星的地球化处理，然后使用该程序对 Mars、Uranus、Neptune 实行地球化
// •必须使用 planets 类型，并之实现相应的 terraform 方法。
type planets []string

func (planet planets) terraform() (turnedPlanet planets) {
	for i, d := range planet {
		planet[i] = "New" + d
	}
	return planet
}
func main() {
	str := planets{"Mars", "Uranus", "Neptune"}
	fmt.Println(str.terraform())
}
