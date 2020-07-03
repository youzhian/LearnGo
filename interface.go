package main

import "fmt"

/*定义接口*/
type Phone interface {
	call()
}

type NokiaPhone struct {
	name string
}

func (phone NokiaPhone) call() {
	fmt.Printf("这个是%s手机\n", "诺基亚")
}

type Iphone struct {
	name string
}

func (phone Iphone) call() {
	fmt.Printf("这个是%s手机\n", "苹果")
}

func main() {

	var phone Phone

	phone = new(NokiaPhone)
	phone.call()
	phone = new(Iphone)
	phone.call()
}
