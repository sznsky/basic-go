package main

//常量声明方式：常量声明不使用，不会报错，不能修改

// 包变量 : 大小写控制访问权限
const External = "包外"
const internal = "保内"

// 包变量-块状声明：大小写控制访问权限
const (
	a = 123
	A = 456
	C = "789"
)

// iota的使用,iota默认是从0开始的,依次次递增的
const (
	StatusA = iota
	StatusB
	StatusC
	d
	e
	hha
	//还可以主动中断,下面这些写都是一样的
	hh = 100
	ff
)

// iota还可以这样使用
const (
	Day1 = iota*12 + 13
	Day2            //这里为什么是25，因为iota是从0开始的，这里隐藏了算法：1*12+13
	Day3 = iota + 1 //这里为什么是3，因为iota是从0开始的，新的算法（iota+1）会替代老的算法(iota*12+13)
)

// iota还可以使用位移，其实iota就是代表的0,1,2,3,4...
const (
	//0左移一位
	Good1 = iota << 1
	//1左移一位
	Good2
	//2左移一位
	Good
)

func main() {

	//局部变量
	const a = 123
	//a = 456 常量不能修改

	const (
		b = 456
		c = "hello go"
	)

	//注意：iota要实现递增，局部变量块也可以使用
	const (
		a1 = iota
		a2
	)

	const (
		haha = iota*1 + 3
		hah1
		haha2 = iota + 10
	)

	//总结：iota要实现递增，只能在包变量块，或者局部变量块使用

}
