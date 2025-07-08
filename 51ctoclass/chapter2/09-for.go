package main

import (
	"fmt"
	"time"
)

func main() {
	// 计算 1+2+3...+100

	//1.第一种方法
	i, sum := 0, 0
	for i = 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	//2.第二种方法
	i = 1
	sum = 0
	for i <= 100 {
		sum += i
		i++
	}
	println(sum)

	//3.每隔一秒打印一个haha
	for {
		fmt.Println("haha")
		time.Sleep(1 * time.Second)
	}
}
