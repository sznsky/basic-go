package main

import "fmt"

func main() {
	var s1 []int
	// 创建一个切片，默认容量和长度的是一样的
	s1 = make([]int, 5)
	printSlice(s1)
	// 这里追加一个元素，原来的容量直接翻倍，长度增加一个长度
	s1 = append(s1, 1)
	printSlice(s1)
	// 还可以增加多个长度
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	s1 = append(s1, 4)
	printSlice(s1)

	s2 := make([]int, 5)
	s2 = append(s2, 1)
	printSlice(s2)

	// 将s1 copy赋值给s2,此时的s2和s1是独立分配空间的
	// 如果s2是通过对s1的[2:4]类似这种赋值的，那就说明s2本身是s1的一部分，修改s2会影响s1
	copy(s2, s1)
	printSlice(s2)
	// 给s2赋值，是不影响s1的
	s2[0] = 1000
	printSlice(s2)
	printSlice(s1)
}

// 打印这个切片的相关信息
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
