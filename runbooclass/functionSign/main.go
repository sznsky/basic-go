package main

import "fmt"

// Operation 定义一个函数类型，也叫函数签名
type Operation func(int, int) int

// Add 和函数签名相同的具体函数
func Add(a, b int) int {
	return a + b
}

// Subtract 和函数签名相同的具体函数
func Subtract(a, b int) int {
	return a - b
}

// Calculate 接收函数，作为通用计算函数
func Calculate(a, b int, op Operation) int {
	return op(a, b)
}

func main() {
	//使用加法函数
	resultAdd := Calculate(1, 2, Add)
	fmt.Printf("resultAdd: %d\n", resultAdd)

	//使用减法函数
	resultSub := Calculate(1, 2, Subtract)
	fmt.Printf("resultSub: %d\n", resultSub)

	// 使用匿名函数
	resultMultiply := Calculate(3, 5, func(a, b int) int {
		return a * b
	})
	fmt.Println("resultMultiply:", resultMultiply)
}
