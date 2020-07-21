package main

import "fmt"

func main() {

	var a float32 = 0.70
	var b float32 = 793.43
	//var b float32 = 800
	var c float32 = 100 * b / (100 - a)

	fmt.Printf("总价为：%f,还差：%f", c, b-c)

}
