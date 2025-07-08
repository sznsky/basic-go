package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.Width + r.Height)
}

// CalculateFunc 定义一个函数类型(函数签名)
type CalculateFunc func() int

// PrintResult 接收方法作为参数的通用函数
func PrintResult(calFunc CalculateFunc) {
	fmt.Println("Result:", calFunc())
}
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	PrintResult(rect.Perimeter)
	PrintResult(rect.Area)
}
