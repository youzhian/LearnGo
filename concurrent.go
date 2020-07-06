/*
	Go 并发
	Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。

	goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

	goroutine 语法格式：

	go 函数名( 参数列表 )
	例如：

	go f(x, y, z)
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	go say("我是线程1")
	say("我是主线程")

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}
