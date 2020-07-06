/*
	练习不定参数函数
*/
package main

import "fmt"

func main() {

	practice()
	bibao()
}

func practice() {
	fmt.Println("=========practice()执行==========")
	var a int = 2
	var b int64 = 34
	var c string = "你好"
	var d float64 = 3.4
	var e float32 = 3.2
	test(a, b, c, d, e)
}

/*不定参数*/
func test(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "的数据类型是int类型")
		case string:
			fmt.Println(arg, "的数据类型是string类型")
		case int64:
			fmt.Println(arg, "的数据类型是int64类型")
		case float64:
			fmt.Println(arg, "的数据类型是float64类型")
		default:
			fmt.Println(arg, "是未知数据类型")
		}
	}
}

/*测试闭包函数*/
func bibao() {
	fmt.Println("=========bibao()执行==========")
	var j int = 5
	a := func() {
		var i int = 10
		//return func(){
		fmt.Printf("i,j: %d,%d\n", i, j)
		//}
	}
	//第一次执行
	a()
	j *= 5
	a()
}
