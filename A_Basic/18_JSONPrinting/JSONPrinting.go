package main

import (
	"encoding/json"
	"fmt"
)

// •编写一个程序；
// 、能够以 JSON 格式打印出 例子21.8中3台 探测器的着陆点。
// •被打印的JSON数据必须包含每个着陆点的名称，并使用例子 21.10中展示的struct 标签特性。
// •请使用 json 包中的 MarshalIndent 函数让打印输出变得更加美观和易读。

type coordinate struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func main() {
	opportunity := coordinate{
		Name: "Opportunity",
		Lat:  -1.9462,
		Long: 354.4734,
	}
	insight := coordinate{
		Name: "InSight",
		Lat:  4.5,
		Long: 135.9,
	}
	spirit := coordinate{
		Name: "Spirit",
		Lat:  -14.5684,
		Long: 175.472636,
	}
	l1, _ := json.MarshalIndent(opportunity, "", "  ")
	l2, _ := json.MarshalIndent(insight, "", "  ")
	l3, _ := json.MarshalIndent(spirit, "", "  ")
	fmt.Println(string(l1))
	fmt.Println(string(l2))
	fmt.Println(string(l3))
}
