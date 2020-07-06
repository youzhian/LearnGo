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
	"math/rand"
	"strings"
	"time"
)

func main() {

	go say("我是线程1")
	say("我是主线程")

	channel := make(chan string)
	var isclose bool = false
	go producter("cat", channel, &isclose)
	go producter("dog", channel, &isclose)

	customer(channel)

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

/**
生产者
channelName chan <- type为只写通道
*/
func producter(s string, channel chan<- string, isclose *bool) {
	var i = 0
	var max = 10
	for {
		//若已关闭
		if *isclose {
			fmt.Println(s, " the channel is close")
			break
		}
		if i >= max {
			channel <- s + ": channel closed"
			//关闭通道
			close(channel)
			*isclose = true
			break
		}
		//将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s:%v", s, rand.Int31())
		i += 1
		//休眠一秒
		time.Sleep(time.Second)
	}

}

/*
消费者
channelName <- chan type 为只读通道
*/
func customer(channel <-chan string) {
	//无限循环
	for {
		// 从通道中取出数据, 此处会阻塞直到信道中返回数据
		message := <-channel
		//message=strings.Trim(message)
		fmt.Println(message)
		//若通道关闭，则跳出循环
		if strings.HasSuffix(message, "channel closed") {
			break
		}
	}
}
