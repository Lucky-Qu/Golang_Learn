package main

import "fmt"

// •编写一个程序
// •它包含三种类型：celsius、fahrenheit、kelvin
// •3 种温度类型之间转换的方法
type celsius float64
type fahrenheit float64
type kelvin float64

func (c celsius) Celsius() (f fahrenheit, k kelvin) {
	f = fahrenheit(9*c/5 + 32)
	k = kelvin(c + 273.15)
	return f, k
}
func (f fahrenheit) Fahrenheit() (c celsius, k kelvin) {
	c = celsius((f - 32) * 5 / 9)
	k = kelvin(c + 273.15)
	return c, k
}
func (k kelvin) Kelvin() (c celsius, f fahrenheit) {
	c = celsius(k - 273.15)
	f, _ = c.Celsius()
	return c, f
}
func main() {
	var t1 celsius = 66
	var t2 fahrenheit = 66
	var t3 kelvin = 66
	t11, t12 := t1.Celsius()
	fmt.Printf("摄氏温度：%.2f华氏温度：%.2f开氏温度：%.2f\n", t1, t11, t12)
	t21, t22 := t2.Fahrenheit()
	fmt.Printf("摄氏温度：%.2f华氏温度：%.2f开氏温度：%.2f\n", t21, t2, t22)
	t31, t32 := t3.Kelvin()
	fmt.Printf("摄氏温度：%.2f华氏温度：%.2f开氏温度：%.2f\n", t31, t32, t3)
}
