package main

import (
	"fmt"
	"sync"
)

func main() {
	// waitGroup就是一个计数器
	var wg sync.WaitGroup
	// 10个goroutine干活
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(num int) {
			// 如果去掉这行，就是无序的
			//time.Sleep(time.Duration(num) * time.Second)
			fmt.Println("I am ", num, "goroutine")
			// 当前的事情干完了，标记一下
			wg.Done()
		}(i)
	}
	// 等待所有的任务做完，退出
	wg.Wait()
	fmt.Println("============================")
}
