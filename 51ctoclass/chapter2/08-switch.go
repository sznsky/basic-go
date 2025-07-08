package main

import "fmt"

func SwitchTest() {
	//定义一个水果名称
	var fruit string
	//等待标准输入,将值赋给fruit
	fmt.Scanf("%s", &fruit)

	switch fruit {
	case "apple":
		fmt.Printf("A,you choose apple")
	case "banana":
		fmt.Printf("B,you choose bananan")
	case "orange":
		fmt.Printf("c,you choose orange")
	default:
		fmt.Printf("D,you choose bananan")
	}
}

func SwitchTest1() {
	//通过年龄判断是什么人
	var age int
	//等待标准输入,将值赋给age
	fmt.Scanf("%s", age)

	switch {
	case age >= 18:
		fmt.Printf("成年人")
	case age >= 12 && age < 18:
		fmt.Printf("青少年")
	default:
		fmt.Printf("儿童")
	}
}
