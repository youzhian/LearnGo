/**
数组
*/
package main

import "fmt"

func main() {
	/*
		数组的声明
		var array_name [size] array_item_type
	*/
	//声明一个长度为10的数组
	var balances [10]float32
	fmt.Println("balances的长度：", len(balances))
	//初始化
	var array1 = [5]float32{10, 4.2, 54.4, 471, 4.1}
	var balance2 = [...]float32{3, 2.4, 54, 34, 3.54, 2}
	fmt.Println("array1的长度：", len(array1))
	fmt.Println("balance2的长度：", len(balance2))
	fmt.Println("赋值前：", array1)
	array1[4] = 33
	fmt.Println("赋值后：", array1)
	fmt.Println(balances)

	//调用函数
	practice()

	pract()
}

func practice() {
	fmt.Println("==========practice方法被执行============")
	//声明一个长度为10的整型数组
	var ls [10]int
	var i, j int
	//循环赋值
	for i = 0; i < 10; i++ {
		ls[i] = i + 100
	}
	//循环遍历
	for j = 0; j < len(ls); j++ {

		fmt.Printf("Element[%d] = %d\n", j, ls[j])
	}
}

func pract() {
	fmt.Println("========执行pract()========")
	var numbers = [3]int{1, 2, 3}

	fmt.Println(numbers)
}
