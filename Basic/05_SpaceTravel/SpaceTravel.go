package main

//实现一个将给定公司和距离的数值，随机生成速度并生成相应旅行速度，随机生成单程还是双程，并给出一个完善的表格
import (
	"fmt"
	"math/rand"
)

const distance = 62100000

func main() {
	company := []string{
		"Space Adventures",
		"SpaceX",
		"Virgin Galactic",
	}
	var tripType = []string{
		"One-Way",
		"Round-Type",
	}
	fmt.Printf("%-16s %4s %16s %8s\n", "Spaceline", "Days", "Trip Type", "Price")
	fmt.Println("------------------------------------------------")
	for i := 0; i < 10; i++ {
		r := rand.Intn(len(tripType))
		trip := tripType[r]
		days := (distance / ((rand.Intn(15) + 16) * 3600 * 24)) * (r + 1)
		fmt.Printf("%-16s %4d %16s %8d\n", company[rand.Intn(3)], days, trip, (r+1)*(rand.Intn(14)+36))
	}
}
