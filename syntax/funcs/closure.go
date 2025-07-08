package main

import "fmt"

// 闭包:返回值就是一个函数,同时函数的返回值是字符串
func Closure(name string) func() string {
	return func() string {
		return "hello," + name
	}
}

// 闭包1
func Closure1() func() int {
	var age = 18
	fmt.Printf("out: %p ", &age)
	return func() int {
		age++
		fmt.Printf("%p ", &age)
		return age
	}
}
