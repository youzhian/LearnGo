package main

import (
	"errors"
	"fmt"
)

/*定义一个结构体*/
type DivideError struct {
	dividee int
	divider int
}

/*
	定义一个书结构体
*/
type BookError struct {
	name   string
	author string
}

/*错误信息处理*/
func main() {

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
	book := BookError{name: "三体", author: "刘慈欣"}
	var msg = book.Error()
	fmt.Println(msg)
}

func Sqrt(f float64) (float64, error) {

	if f < 0 {
		return 0, errors.New("f不能小于0")
	}
	return f, nil
}

/*实现 error接口*/
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

/*实现 error接口*/
func (book *BookError) Error() string {
	return "《" + book.name + "》|" + book.author
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}
