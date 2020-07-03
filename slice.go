//切片
package main

import "fmt"

func main() {
	/*
		数组和切片的区别在于声明或者初始化时数组需要指定长度，而且长度不可变，
		而切片则不需要指定长度，且长度可变
		通过len()可计算长度，cap()可获取容量
		make([]T,len,cap)，cap是可选参数
	*/
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	pract()

	pract2()

	pract1()

	arrayAndSlice()
}

func printSlice(x []int) {

	fmt.Printf("len=%d cap=%d slic=%v\n", len(x), cap(x), x)
}

func pract() {
	fmt.Println("========pract()========")
	//空切片
	var numbers []int

	printSlice(numbers)

	if numbers == nil {
		fmt.Println("numbers是空切片")
	}
}

func pract2() {
	fmt.Println("==========pract2()=========")
	//初始化切片变量
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	fmt.Println("numbers ==", numbers)
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
}

/**
使用append与copy方法对数组扩容
*/
func pract1() {
	fmt.Println("========pract1()========")
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	getAddr(numbers)
	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)
	getAddr(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func getAddr(slice []int) {
	fmt.Println("=============执行getAddr()===========")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d]=%d,地址为：0x%x\n", i, slice[i], &slice[i])
	}
}

/**
切片与数组的区别
当append时的内容大于原来切片的cap，将不会再指向旧地址
*/
func arrayAndSlice() {
	fmt.Println("=========执行arrayAndSlice()========")
	var x = [3]int{2, 4, 6}
	var y = x
	y[1] = 200
	fmt.Printf("x=%d，y=%d\n", x, y)

	var x1 = []int{2, 4, 6}
	var y1 = x1
	y1[1] = 300
	fmt.Printf("x1=%d，y1=%d\n", x1, y1)

	var x2 = make([]int, 0, 5)
	x2 = append(x2, 1, 9, 5)
	fmt.Printf("x2 = %d\n", x2)
	var y2 = append(x2, 3, 6)
	y2[1] = 504
	fmt.Printf("x2=%d，y2=%d\n", x2, y2)
}
