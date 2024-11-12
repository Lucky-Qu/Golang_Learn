package main

// 实现维吉尼亚算法加密和解密一个信息
func encryption(str, password string) (result string) {
	var mid = []rune(str)
	var pass = []rune(password)
	for i := range mid {
		mid[i] += pass[i%len(pass)] - 'A'
		if mid[i] > 'Z' {
			mid[i] -= 26
		}
	}
	result = string(mid)
	return result
}
func decrypt(str, password string) (result string) {
	var mid = []rune(str)
	var pass = []rune(password)
	for i := range mid {
		mid[i] -= pass[i%len(pass)] - 'A'
		if mid[i] < 'A' {
			mid[i] += 26
		}
	}
	result = string(mid)
	return result
}
func main() {
	str := "TESTTEST"
	password := "WTF"
	str = encryption(str, password)
	println(str)
	str = decrypt(str, password)
	println(str)
}
