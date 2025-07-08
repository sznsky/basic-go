package main

import (
	"fmt"
	"sync"
)

var x int = 0
var mutex sync.Mutex

// 加锁过程：1.请求前加锁，2修改数据 3-解锁
func increment() {
	mutex.Lock()
	// 这个x就是共享资源，如果不加锁，使用goroutine的时候，会有很多协程进来
	x += 1
	mutex.Unlock()
	wg.Done()
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	fmt.Println("final value is ", x)
}
