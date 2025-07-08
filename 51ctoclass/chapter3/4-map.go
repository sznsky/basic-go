package main

import "fmt"

func main() {
	// 构造一个map, key和value分别是string,string
	countryCapitalMap := make(map[string]string)

	// 给map的key赋值
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	countryCapitalMap["China"] = ""

	//直接打印
	fmt.Println(countryCapitalMap)
	//通过map获取中的Key获取val的值
	val := countryCapitalMap["France"]
	fmt.Println(val)
	val = countryCapitalMap["China"]
	fmt.Println(val)
	// map指示器的使用
	val, ok := countryCapitalMap["China"]
	if ok {
		fmt.Println("China capital is not exist")
	} else {
		fmt.Println("China capital is ", val)
	}

}
