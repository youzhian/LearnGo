/*
	通道（channel）
	通道（channel）是用来传递数据的一个数据结构。

	通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。

	ch <- v    // 把 v 发送到通道 ch
	v := <-ch  // 从 ch 接收数据,并把值赋给 v

	声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：

	ch := make(chan int)
	注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。

	以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：
*/
/*
有关 channel 的关闭，你需要注意以下事项:

关闭一个未初始化(nil) 的 channel 会产生 panic
重复关闭同一个 channel 会产生 panic
向一个已关闭的 channel 中发送消息会产生 panic
从已关闭的 channel 读取消息不会产生 panic，且能读出 channel 中还未被读取的消息，若消息均已读出，则会读到类型的零值。从一个已关闭的 channel 中读取消息永远不会阻塞，
并且会返回一个为 false 的 ok-idiom，可以用它来判断 channel 是否关闭
关闭 channel 会产生一个广播机制，所有向 channel 读取消息的 goroutine 都会收到消息
*/
package main

import "fmt"

func main() {

	practic()
	practic2()
	practic3()
}

/*通过通道传值*/
func practic() {
	fmt.Println("========practic()被执行=======")
	var s = []int{1, 5, 7, 3, 0, -3, 9, 55}
	var c = make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

/*若无缓冲区着写入后必须进一步获取，否则下一次写入会报错*/
func practic2() {
	fmt.Println("========practic2()被执行=======")
	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	ch := make(chan int, 2)
	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2
	x, y := <-ch, <-ch
	// 获取这两个数据
	fmt.Println(x)
	fmt.Println(y)
	/*
		若未设置缓冲区，则会在第二次写入时报错，除非用goroutine去执行
		ch := make(chan int)
		ch <- 1
		ch <- 2
		x,y := <-ch,<-ch
	*/
}

/*
Go 遍历通道与关闭通道
Go 通过 range 关键字来实现遍历读取到的数据，类似于与数组或切片
*/
func practic3() {
	fmt.Println("========practic3()被执行=======")
	//定义一个缓冲区为10的通道
	c := make(chan int, 10)
	//var n int = 11
	fibonacci(cap(c), c)
	for v := range c {
		fmt.Println(v)
	}
}

func sum(s []int, c chan int) {
	var sum int
	for _, v := range s {
		sum += v
	}
	//将sum的值发送到通道中
	c <- sum
}

/*计算第n项的斐波那契数，并返回其前面的斐波那契数*/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
