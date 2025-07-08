package main

import "fmt"

func main() {
	// 定义5个元素的数组-整型
	var five [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(five)

	// 定义3个元素的数组-字符串
	var three [3]string = [3]string{"lilei", "lucy", "hanmei"}
	fmt.Println(three)
	// 这个也可以这样写，初始化的可以少些,但是不能多个实例，因为数组是固定长度的
	var three1 [3]string = [3]string{"lilei"}
	fmt.Println(three1)

	// 上面的都可以这样命名
	three2 := [3]string{"lilei", "lucy", "hanmei"}
	fmt.Println(three2)
}
