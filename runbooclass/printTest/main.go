package main

import "fmt"

func main() {

	//函数的调用
	x := Add(1, 2)
	fmt.Println(x)

	//给结构体赋值
	p := Person{name: "sunzhenning", age: 18}
	//打印对象的属性值
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	//方法的调用
	fmt.Println("Name:", p.Name())
	fmt.Println("Age:", p.Age())

	//打印多个值
	fmt.Println("Name:", p.name, "Age:", p.age)

	//

}
