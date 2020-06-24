//定义包路径，若为main表示该文件可独立执行的程序
package main
//引入fmt
import "fmt"
//程序入口，若有init函数则先执行init()
func main()  {
	/**
	fmt.Println("Test Two!")
	 */
	//字符串相加
	fmt.Println("Goolge"+ "Runoob")
	// 若变量名首字母大写，则对外部可见，类似于java的public，否则只是整个包可见
	var age int =23

	fmt.Println("my age is :" ,age)
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
	const (
		a1 = "aa"
		b2 = "bb"
		c2
	)
	fmt.Println(a1,b2,c2)

	const (
		i1 = 1 << iota
		j1 = 3 << iota
		k
		l
	)
	fmt.Println(i1,j1,k,l)
	var a2 bool = false
	var b1 bool = true

	if a2 && b1{
		fmt.Println("第一行条件为true")
	}
	if a2 || b1 {
		fmt.Println("第二行条件为true")
	}
	a2 = true
	b1 = true
	if a2 && b1{
		fmt.Println("第三行条件为true")
	}
}
