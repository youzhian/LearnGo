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
}
