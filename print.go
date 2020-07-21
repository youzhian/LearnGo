package main

import "fmt"

func main() {
	fmt.Println("开始进行星号输出")
	var len = 10
	var p string = "1"
	for n := 1; n <= len; n++ {
		x := getX(n, len, p)
		fmt.Println(x)
	}
	//打印圣诞树的树干
	for i := 0; i < len/3; i++ {
		x := getX(len/3, len, p)
		fmt.Println(x)
	}
}

func getX(n int, lenNum int, prifex string) string {
	x := ""
	if n < 1 {
		return x
	}
	p := "*"
	if len(prifex) > 0 {
		p = prifex
	}
	num := 2*n - 1
	for i := 0; i < num; i++ {
		x = x + p
	}
	pl := len(p)
	//拼凑空格
	max := 2 * (lenNum - n) * pl
	for i := 0; i < max/2; i++ {
		x = " " + x
	}
	return x
}
