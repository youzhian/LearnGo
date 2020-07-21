package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Lstat("F:/bootstrap.yml")
	if err != nil {
		fmt.Println("读取错误：", err)
		return
	}
	fmt.Println("文件名：", fileInfo.Name())
	fmt.Println("文件大小：", fileInfo.Size())
	fmt.Println("是否为目录：", fileInfo.IsDir())
	var a float64 = 480 / (8 * 12)
	fmt.Println(a)
}
