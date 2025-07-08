package main

import (
	"fmt"
	"sync"
)

func main() {
	// 共享数据
	var count int = 0
	//
	var wg sync.WaitGroup
	var once sync.Once

	// 增加100个goroutine
	wg.Add(100)

	for i := 0; i < 100; i++ {
		// 匿名函数
		go func() {
			// once do 里面也是一个匿名函数，这个函数只是执行一次
			once.Do(func() {
				count += 1
			})
			wg.Done()
		}()
	}
	// 等待所有的函数结束:等待所有的goroutine结束
	wg.Wait()
	fmt.Printf("count = %d\n", count)
}
