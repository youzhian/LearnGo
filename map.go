/*
	集合
*/
package main

import "fmt"

func main() {

	/*创建空集合*/
	var countryCapitalMap map[string]string
	/*使用make创建集合*/
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "的首都是：", countryCapitalMap[country])
	}

	//判断某个国家的首都是否在集合中
	capital, ok := countryCapitalMap["American "]
	if ok {
		fmt.Println("American 的首都是：", capital)
	} else {
		fmt.Println("未找到American 的首都")
	}
	practDelete()
}

/**
练习删除集合
*/
func practDelete() {
	fmt.Println("=========practDelete()=======")
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")
	//遍历集合
	for country, capital := range countryCapitalMap {
		fmt.Println(country, "的首都是", capital)
	}
	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国被删除")
	fmt.Println("删除后的地图")

	for country, capital := range countryCapitalMap {
		fmt.Println(country, "的首都是", capital)
	}
}
