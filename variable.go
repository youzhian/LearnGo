package main

import (
	"fmt"
)

func main() {
	var a int = 4
	var b float32
	var c float64
	var prt *int

	fmt.Printf("第一行 a的类型为%T\n", a)
	fmt.Printf("变量b的类型为%T\n", b)
	fmt.Printf("变量c的类型为%T\n", c)

	prt = &a
	fmt.Printf("a的值为：%d\n", a)
	fmt.Printf("prt的值为：%d\n", prt)
	fmt.Printf("*prt的值为：%d\n", *prt)

	test()
}

func test() {
	var a int = 4
	var prt *int = &a
	fmt.Printf("赋值前：a: %d\n", a)
	fmt.Printf("赋值前：&a: %x\n", &a)
	fmt.Printf("赋值前：*prt: %d\n", *prt)
	fmt.Printf("赋值前：prt: %x\n", prt)

	*prt = 5
	fmt.Printf("赋值后：a: %d\n", a)
	fmt.Printf("赋值后：&a: %x\n", &a)
	fmt.Printf("赋值后：*prt: %d\n", *prt)
	fmt.Printf("赋值后：prt: %x\n", prt)
}
