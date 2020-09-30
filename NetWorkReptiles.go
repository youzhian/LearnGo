//网络爬虫，并发版
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const url = "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn="

//get请求
func HttpGet(url string) (result string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//程序运行结束后关闭
	defer response.Body.Close()
	buf := make([]byte, 1024*6)
	for {
		n, _ := response.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf)
	}
	return
}

//将文件保持到本地，c为一个可读写的通道
func SaveFileToLocal(index int, c chan<- int) {

	indexStr := strconv.Itoa((index - 1) * 50)
	//网络请求获取内容
	result := HttpGet(url + indexStr)

	if result != "" {
		fmt.Println("the result is not empty")
		//获取当前绝对路径
		path, _ := os.Getwd()
		f, err := os.Create(path + "/page/" + strconv.Itoa(index) + ".html")
		if err != nil {
			fmt.Println(err)
		}
		//将内容写到文件
		f.WriteString(result)
		f.Close()
	} else {
		fmt.Println("the result is empty !")
	}
	//告诉协程当前请求的页数
	c <- index
}

func SaveFile(index int) {
	indexStr := strconv.Itoa((index - 1) * 50)
	//网络请求获取内容
	result := HttpGet(url + indexStr)

	if result != "" {
		//获取当前绝对路径
		path, _ := os.Getwd()
		f, err := os.Create(path + "/page/" + strconv.Itoa(index) + ".html")
		if err != nil {
			fmt.Println(err)
			return
		}
		//将内容写到文件
		f.WriteString(result)
		f.Close()
	}
}

//爬取网页
func doWork(start, end int) {
	//创建无缓冲通道
	page := make(chan int)
	for i := start; i <= end; i++ {
		//开启协程
		go SaveFileToLocal(i, page)
	}

	for i := start; i <= end; i++ {
		//等待协程返回
		fmt.Printf("第%d页被爬取完毕\n", <-page)
	}
	//关闭通道
	close(page)
}

func main() {
	var start, end int
	start = 1
	end = 10
	doWork(start, end)

}
