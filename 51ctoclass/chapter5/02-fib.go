package main

import (
	"fmt"
	"time"
)

func main() {
	// 启动了一个观感服务
	go spinner(time.Millisecond * 100)
	// 时间很长
	fmt.Println("\n", fib(45))
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

// 增加用户体验
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
