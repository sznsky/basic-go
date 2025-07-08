package main

import "fmt"

type Person struct {
	name  string
	age   int
	sex   string
	fight int
}

func main() {
	p := Person{
		name:  "lily",
		age:   22,
		sex:   "女",
		fight: 1,
	}
	fmt.Println(p)

	//如果这种写法，需要全部赋值
	p1 := Person{"lilei", 18, "男", 100}
	fmt.Println(p1)

}
