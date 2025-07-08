package main

import "fmt"

// 数组
func Array() {
	//声明一个数组
	a1 := [3]int{1, 5, 8}
	//%v直接打印整个数组，len求出长度的方法，cap求出容量的方法
	fmt.Printf("a1: %v, len=%d, cap=%d \n", a1, len(a1), cap(a1))

	//声明一个数组，但是只给了2个值，长度是3，自动补0
	a2 := [3]int{1, 3}
	fmt.Printf("a1: %v, len=%d, cap=%d \n", a2, len(a2), cap(a2))

	//声明一个数组，不赋值，会自动补0
	var a3 [3]int
	fmt.Printf("a1: %v, len=%d, cap=%d \n", a3, len(a3), cap(a3))

}
