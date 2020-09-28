package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//定义结构体
type Handler struct {
}

//首页
func (h *Handler) Index(writer http.ResponseWriter, request *http.Request) {

	var data struct {
		Title string
		Body  string
	}
	data.Title = "Go Web"
	data.Body = "welcome!"
	//绘制页面
	t, err := template.ParseFiles("web/page/index.html")
	if err != nil {
		fmt.Println(writer, err)
		return
	}
	t.Execute(writer, data)

}

//登录接口
func (h *Handler) Login(writer http.ResponseWriter, request *http.Request) {

	if err := request.ParseForm(); err != nil {
		fmt.Println(writer, err)
		return
	}
	//获取前端传递的用户名和密码
	username := request.FormValue("username")
	password := request.FormValue("password")
	if username == "admin" && password == "admin" {
		writer.Write([]byte("登录成功"))
	} else {
		writer.Write([]byte("用户名密码不正确"))
	}

}

func main() {
	//初始化结构体
	handler := &Handler{}
	//设置过滤器
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/login", handler.Login)
	//开启监听并启动服务
	http.ListenAndServe(":8080", nil)

}
