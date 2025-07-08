package main

import (
	"fmt"
	"time"
)

func main() {
	// 主函数内置的一个goroutine
	fmt.Println("I am main goroutine")
	// 这是一个自定义的goroutine，异步的
	go func() {
		fmt.Println("I am a goroutine")
	}()
	fmt.Println("bye bye")
	// main函数退出，进程退出，所有的goroutine也就是没有任何意义,sleep意义确保其实的goroutine运行完成
	time.Sleep(time.Second * 2)
}
