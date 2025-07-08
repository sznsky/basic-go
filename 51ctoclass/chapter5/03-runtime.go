package main

import (
	"fmt"
	"runtime"
)

func main() {
	// GOMAXPROCS 设置是小于1的时候，是查看当前最大cpu的核心数，大于等于1的时候是设置
	fmt.Println("max proc is ", runtime.GOMAXPROCS(1))
	runtime.GOMAXPROCS(3)
	fmt.Println("max proc is ", runtime.GOMAXPROCS(1))
	// 释放cpu
	runtime.Gosched()

}
