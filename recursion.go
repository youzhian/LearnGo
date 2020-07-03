/*递归*/
package main

import "fmt"

/*
递归练习
递归，就是在运行的过程中调用自己
*/
func main() {

	var i uint64 = 15
	fmt.Printf("%d 的阶乘为：%d\n", i, factorial(i))
	fmt.Printf("%d 的阶乘为：%d\n", i, factorial2(i, 0))
	fmt.Println("执行斐波那契数")
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println()
	fmt.Println("执行尾递归斐波那契数")
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacci2(i, 0, 1), " ")
	}
	fmt.Println()
	//fmt.Println(fibonacci2(4,0))
	changeType()
}

/*
实现求阶乘
n 为参数，即要进行求阶乘的数
result 为返回值
*/
func factorial(n uint64) uint64 {

	if n > 0 {
		return n * factorial(n-1)
	}
	return 1
}

/*尾递归实现阶乘*/
func factorial2(n uint64, sum uint64) uint64 {
	if sum < 1 {
		sum = 1
	}
	if n == 0 {
		return sum
	}
	return factorial2(n-1, sum*n)
}

/**
  斐波那契数
*/
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

/**
尾递归
*/
func fibonacci2(n int, first int, second int) int {

	if n < 1 {
		return first
	}

	return fibonacci2(n-1, second, first+second)
}

func changeType() {
	fmt.Println("===========s()===========")
	var a = 15
	var b = 4
	//类型转换
	fmt.Printf("%f", float32(a/b))
}
