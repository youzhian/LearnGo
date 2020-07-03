/*
	结构体
*/
package main

import "fmt"

/**
定义书的结构体
*/
type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {

	pract()

	pract2()
}

func pract() {
	// 创建一个新的结构体
	fmt.Println(Book{"GoLang从入门到入狱", "youzhian", "golang语言是一门解析性语言", 31})
	// 用键值的方式创建结构体
	fmt.Println(Book{title: "Golang语言", subject: "golang语言是又google公司开发的", author: "游志安", bookId: 232})
	// 省略部分值
	fmt.Println(Book{title: "Golang语言", subject: "golang语言是又google公司开发的", bookId: 345})
}

func pract2() {
	var book1 Book
	var book2 Book

	book1.title = "golang编程语言"
	book1.bookId = 34234
	book1.author = "张三"
	book2.title = "三体-死神永生"
	book2.bookId = 345345

	fmt.Println("第一本书的标题为:", book1.title)
	fmt.Println(book1)
	fmt.Println("第二本书的标题为:", book2.title)
	fmt.Println(book2)
}
