package main

import "fmt"

func main() {

	// 定义一个二维数组
	a2 := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(a2)

	// 遍历这个二维数组
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(a2[i][j])
		}
	}
}
