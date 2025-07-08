package main

import "strings"

//方法简介

// 1.无返参数，无返回值
func Fun1() {

}

// 2.有参数，无返回值
// 一个参数
func Fun2(a int) {

}

// 两个参数
func Fun3(a int, b int) {

}

// 同一个类型参数，可以合并写
func Fun4(a, b int) {

}

// 不同类型的参数
func Fun5(a int, b string) {

}

// 3.有返回值,一个
func Fun6(a, b string) string {
	return "hello world"
}

// 有多个返回值
func Fun7(a, b string) (string, string) {
	return "zhangsan", "lisi"
}

// 返回值可以带名称
func Fun8() (name string, age int) {
	return "lucy", 18
}

// 上面的还可以这样写,赋值再返回
func Fun9() (name string, age int) {
	name = "lucy"
	age = 18
	return
}

// 如果不赋值，直接return
func Fun10() (name string, age int) {
	//这里等价于 return "",0 相当于给每个类型赋值，初始类型
	return
}

// 上面的还可以这样写,等价这种
func Fun11() (string, int) {
	var name string
	var age int
	return name, age
}

//下面这种写法是不行的，返回值，要么都给名字，要么都不给名字
//func Fun12() (name string, int) {
//	return
//}

// 参数，返回值string，int 但是没有名字
func Func13(abc string) (string, int) {
	//将字符串abc按照" ",分割，得到字符串数组
	segs := strings.Split(abc, " ")
	//返回字符串数组第一个元素，和整个数组的长度
	return segs[0], len(segs)
}

// 如果给返回值增加名称，如果中途不使用，不建议添加名字,这样无形扩大了作用域
func Func14(abc string) (first string, length int) {
	segs := strings.Split(abc, " ")
	first = segs[0]
	length = len(segs)
	return first, length
}
