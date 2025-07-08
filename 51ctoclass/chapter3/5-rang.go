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

	for k, v := range countryCapitalMap {
		fmt.Println(k, v)
	}

	// range 也可以遍历数组
	five := [5]int{1, 2, 3, 4, 5}
	for k, v := range five {
		fmt.Printf("a[%d] = %d\n", k, v)
	}
	for _, value := range five {
		fmt.Println(value)
	}
}
