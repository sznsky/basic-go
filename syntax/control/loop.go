package main

import "fmt"

func ForLoop() {

	//标准的for循环写法
	for i := 0; i < 10; i++ {
		println(i)
	}

	//拓展写法1.i++还可以写到程序里面
	for i := 0; i < 10; {
		println(i)
		i++
	}

	//拓展写法2，i的声明还可以写到程序外面
	i := 0
	for ; i < 10; i++ {
		println(i)
	}

	//拓展写法3
	for i < 10 {
		println(i)
		i++
	}

	//死循环1
	for true {
		i++
		println(i)
	}

	//死循环2
	for {
		i++
		println(i)
	}
}

// 遍历数组
func ForArr() {
	arr := [3]int{1, 2, 3}
	//可以这样遍历1
	for index, value := range arr {
		println("下标：", index, " 值：", value)
	}

	println("------------------")

	//遍历2
	for index := range arr {
		println("下标：", index, " 值：", arr[index])
	}
}

// 遍历切片：切片说白了，就是不定数组
func ForSlice() {
	slice := []int{1, 2, 3}
	//可以这样遍历1
	for index, value := range slice {
		println("下标：", index, " 值：", value)
	}

	println("------------------")

	//遍历2
	for index := range slice {
		println("下标：", index, " 值：", slice[index])
	}
}

func ForMap() {
	//定义一个map类型的变量m
	m := map[string]int{
		"key1": 100,
		"key2": 200,
	}

	//变量1,变量k和v
	for k, v := range m {
		if v == 200 {
			break
		}
		println(k, v)
	}

	//遍历2，遍历k
	for k := range m {
		println(k, m[k])
	}

}

func ForBug() {
	//定义一个切片类型的变量users,里面存放的是User
	users := []User{
		{
			name: "大明",
		},
		{
			name: "小明",
		},
	}
	//定义一个变量m,是map类型的，键值对分别是string和*User类型（user的指针），使用make方法创建
	m := make(map[string]*User)
	//变量users，同时给map赋值
	for _, u := range users {
		fmt.Printf("%p \n", &u)
		m[u.name] = &u
	}
	//打印这个m变量
	fmt.Printf("%v", m)

	//注意：不要对迭代地址取地址

}

type User struct {
	name string
}

// for循环的终止
func LoopBreak() {
	i := 0
	for true {
		if i > 10 {
			break
		}
		i++
	}
}

// for循环的继续执行
func LoopContinue() {
	for i := 0; i < 10; i++ {
		//如果是偶数，继续下个循环
		if i%2 == 0 {
			continue
		}
		println(i)
	}
}
