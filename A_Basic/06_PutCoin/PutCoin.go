package main

import (
	"fmt"
	"math/rand"
)

// 随机的向罐子中投入0.05,0.10和0.25的硬币，直到有20美元
func main() {

	coin := []float64{
		0.05,
		0.10,
		0.25,
	}
	for money := 0.0; money < 20; {
		money += coin[rand.Intn(len(coin))]
		fmt.Printf("目前存了有%.2f\n", money)
	}
}
