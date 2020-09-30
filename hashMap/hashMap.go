package main

import "fmt"

//定义hashMap结构体
type HashMap struct {
	//node Node
}

//定义切片
var table [](*HashMap)

//切片初始默认长度
var defLength = 10

func main() {

	getHashCode2("中国")
	getHashCode2("美国")
	getHashCode2("abcd")
	getHashCode2("abcde")
	getHashCode2("abcdef")
	getHashCode2("abcdefg")
	getHashCode2("中国人民银行")
	getHashCode2("中国银行")
}
func getHashCode2(key string) {
	fmt.Printf("'%s' 对应的hashCode为：%d\n", key, genHashCode(key))
}

//生成hashCode
func genHashCode(k string) int {
	if len(k) == 0 {
		return 0
	}
	hashCode := 0
	lastIndex := len(k) - 1
	for i := range k {
		if i == lastIndex {
			hashCode += int(k[i])
			break
		}
		hashCode += hashCode*31 + int(k[i])
	}
	return hashCode
}

//初始化
func init() {
	if len(table) == 0 {
		for i := 1; i < defLength; i++ {
			//table[i] = &HashMap{"","",i,nil}
		}
	}
}

func New() *HashMap {

	return &HashMap{}
}
