//指针案例
package main

import "fmt"

func main() {
	//声明变量
	var a int = 10
	//声明指针变量
	var ip *int
	//给指针变量赋值
	ip = &a

	fmt.Println("a的内存地址为：", &a)

	fmt.Printf("ip变量存储的指针地址：%x\n", ip)
	fmt.Printf("ip变量的值：%x\n", *ip)

	//调用函数
	practice()
	//调用函数
	practice2()
}

/**
 * 练习
 */
func practice() {
	//空指针
	var ptr *int
	fmt.Printf("prt的值为：%x\n", ptr)
	if ptr != nil {
		fmt.Println("prt不是空指针")
	}
	if ptr == nil {
		fmt.Println("prt是空指针")
	}
}

/**
 * 指向指针的指针
 */
func practice2() {

	var a int = 3000

	var ptr *int

	var pptr **int
	//指针 ptr地址
	ptr = &a
	//指向指针ptr地址
	pptr = &ptr

	fmt.Printf("变量a的值%d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}
