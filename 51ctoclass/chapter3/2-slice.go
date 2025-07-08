package main

import "fmt"

func main() {

	// 定义一个数组
	a1 := [5]int{1, 2, 3, 4, 5}
	// 定义一个切片(其实就是动态数组)（后面的起始index,结束index,含头不含尾）
	s1 := a1[2:4]
	fmt.Println(s1)

	s1[1] = 100
	fmt.Println(s1)
	fmt.Println(a1)
	//注意：切片和原数组都是被修改的，因为这个切片其实就是原数组的一部分。本质就是修改的原数组
}
