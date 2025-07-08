package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var rwLock sync.RWMutex

// 共享数据
var y int = 0

// 读写锁
func main() {
	// 设置10个携程过来干活
	wg.Add(10)

	// 启动7个reader过来干活
	for i := 0; i < 7; i++ {
		go reader(i)
	}
	// 启动3个人writer过来干活
	for i := 0; i < 3; i++ {
		go writer(i)
	}
	wg.Wait()

}

func reader(num int) {
	for i := 0; i < 10; i++ {
		// 读锁
		rwLock.RLock()
		fmt.Printf("I am %d reader goroutine x = %d\n", num, y)
		// 释放读锁
		rwLock.RUnlock()
	}
	wg.Done()
}

func writer(num int) {
	for i := 0; i < 3; i++ {
		// 写锁
		rwLock.Lock()
		y += 1
		fmt.Printf("I am %d writer goroutine x = %d\n", num, y)
		time.Sleep(time.Millisecond * 2)
		// 释放写锁
		rwLock.Unlock()
	}
	wg.Done()
}
