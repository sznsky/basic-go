package main

import "fmt"

func main() {
	// 获得函数指针
	nextNumber := getSequence()
	// 调用函数还是要加上()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	//下面再次调用，和上面的调用的i其实不一样的，各自分配空间的i
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

}

// 这就是一个闭包，返回的是一个匿名函数，这个匿名函数读取了原函数的变量
func getSequence() func() int {
	i := 0
	return func() int {
		i = i + 1
		return i
	}
}
