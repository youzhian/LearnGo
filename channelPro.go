package main

import "fmt"

/*
有关 channel 的关闭，你需要注意以下事项:

关闭一个未初始化(nil) 的 channel 会产生 panic
重复关闭同一个 channel 会产生 panic
向一个已关闭的 channel 中发送消息会产生 panic
从已关闭的 channel 读取消息不会产生 panic，且能读出 channel 中还未被读取的消息，若消息均已读出，则会读到类型的零值。从一个已关闭的 channel 中读取消息永远不会阻塞，
并且会返回一个为 false 的 ok-idiom，可以用它来判断 channel 是否关闭
关闭 channel 会产生一个广播机制，所有向 channel 读取消息的 goroutine 都会收到消息
*/
func main() {
	c := make(chan string)
	go send("hello", c)
	//go read(c)
	fmt.Println("消息：", <-c)
	fmt.Println("消息：", <-c)
	m, f := <-c
	if m == "" {
		fmt.Println("f：", f)
	}
}

/**
定义一个可写通道
*/
func send(msg string, c chan<- string) {
	fmt.Println("send方法被执行")
	if msg != "" {
		//向通道写内容
		c <- msg
	} else {
		c <- "对方给你发了一条空消息"
	}
	//关闭通道
	//close(c)
}

//定义一个只读通道
func read(c <-chan string) {
	fmt.Println("read方法被执行")
	fmt.Println("收到了一条消息：", <-c)
}
