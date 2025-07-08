package main

import (
	"fmt"
)

func main() {

	// 常规玩法
	x, y := 10, 20
	sum, sub := addSub(x, y)
	fmt.Println(sum, sub)

	// 还可以将函数赋值给参数，也就是函数指针,效果是一样的
	funcptr := addSub
	sum1, sub1 := funcptr(x, y)
	fmt.Println(sum1, sub1)

	//下面的玩法注意：
	sum2 := count(100, 10, addFunc)
	sub2 := count(100, 10, subFunc)
	fmt.Println(sum2, sub2)

}

// 函数1，参数两个，返回值也是两个
func addSub(a, b int) (int, int) {
	return a + b, a - b
}

// 函数2, 参数也存在函数：count函数存在a,b参数，还有一个参数math，但是math也是函数
func count(a, b int, math func(a, b int) int) int {
	return math(a, b)
}

// 加法函数： 和上面的math的参数类型，个数，返回值都是一样的
func addFunc(a, b int) int {
	return a + b
}

// 减法函数：和上面的math的参数类型，个数，返回值都是一样的
func subFunc(a, b int) int {
	return a - b
}
