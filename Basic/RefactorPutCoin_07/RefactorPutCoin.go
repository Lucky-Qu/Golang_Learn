package main

import (
	"fmt"
	"math/rand"
)

func main() {
	coin := []int{
		5,
		10,
		25,
	}
	for money := 0; money < 2000; {
		money += coin[rand.Intn(3)]
		fmt.Printf("目前有%d.%0.2d$\n", money/100, money%100)
	}
}
