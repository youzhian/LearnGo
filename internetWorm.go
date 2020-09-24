/*网络爬虫*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("开始进入爬虫程序。。。")
	url := "http://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取网站内容异常", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取内容异常", err)
		return
	}
	fmt.Println(string(body))
	download(url)
}

/*
	更好的读取网页的写法
*/
func download(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//设置用户代理，模拟浏览器访问
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get url error", err)
		return
	}
	//函数结束后执行关闭操作，关闭相关链接
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	fmt.Println(string(body))
}
