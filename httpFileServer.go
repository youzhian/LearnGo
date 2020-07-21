/*
搭建一个简易的文件服务器
*/
package main

import "net/http"

func main() {
	//使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器
	http.Handle("/", http.FileServer(http.Dir("F:/")))
	//开启服务监听，设置端口为8080
	http.ListenAndServe(":8080", nil)
}
