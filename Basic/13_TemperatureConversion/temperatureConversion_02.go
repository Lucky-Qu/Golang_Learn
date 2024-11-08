package main

//实现一个根据温度打印表格的程序

import (
	"fmt"
	"strconv"
)

type fahrenheit float64
type celsius float64

func (f fahrenheit) fahrenheitToCelsius() celsius {
	return celsius((f - 32) * 5 / 9)
}

func (c celsius) celsiusToFahrenheit() fahrenheit {
	return fahrenheit(9*c/5 + 32)
}

func createData(kind string, limit, interval int) (data []string) {
	if data = make([]string, 1); kind == "摄氏度" {
		data[0] = "°C"
	} else if kind == "华氏度" {
		data[0] = "℉"
	}
	for i := 0; -40+i <= limit; i += interval {
		data = append(data, strconv.Itoa(-40+i))
	}
	return data
}
func drawTable(data []string, convert func(celsius) fahrenheit) {
	//打印行
	var line = func() {
		fmt.Println("===================")
	}
	//打印表头
	line()
	if data[0] == "°C" {
		fmt.Printf("|%-8s|%-8s|\n", "°C", "℉")
	}
	line()
	for index := range data {
		if index == 0 {
			continue
		}
		data1, _ := strconv.Atoi(data[index])
		fmt.Printf("|%-8s|%-8.2f|\n", data[index], convert(celsius(data1)))
	}
	line()
}
func main() {
	//生成温度
	data := createData("摄氏度", 40, 5)
	//绘制表格
	drawTable(data, celsius.celsiusToFahrenheit)
}
