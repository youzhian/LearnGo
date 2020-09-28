package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1
	fmt.Println("欢迎来到猜数字游戏，数字的范围是1到100的数。。")
	a := 0
	i := 0
	for {
		if i > 100 {
			fmt.Println("你已经猜了100次了，你不适合玩这个游戏，就此结束吧")
			break
		}
		fmt.Print("请输入你猜的数字，或输入-1退出游戏：")
		fmt.Scan(&a)
		fmt.Println()
		i++
		if a == -1 {
			fmt.Println("看来你不想玩了呢，再见了哦。。")
			break
		}
		if a == num {
			fmt.Println("恭喜你，第", i, "次就答对了")
			break
		} else if a < num {
			fmt.Println("你猜测的数字是：", a, "猜小了，继续加油")
		} else {
			fmt.Println("你猜测的数字是：", a, "猜大了，继续加油")
		}
	}
}
