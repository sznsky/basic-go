package main

// 函数式编程：函数是一等公民
func Functional4() {
	println("hello sunzhenning! functional4")
}

func Functional5(a int) {

}

// 还可以这样写,一个返回值string,不需要加括号
var Abc = func() string {
	return "hello"
}

func UseFunctional4() {
	//直接可以将一个函数赋值给一个变量
	myFunc := Functional4
	//直接对这个变量发起调用
	myFunc()

	//有参数的情况下
	myFunc_1 := Functional5
	myFunc_1(1)
}

// 局部方法,这样也是可以发起调用的：在方法内部封装一个方法，也就是一块逻辑，但是不想给别人使用
func Functional6() {
	//新定义了一个方法（相当于java里面的匿名方法），赋值给了fn
	//注意：局部变量一般不要大写方法名，fn不要写成了Fn
	fn := func() string {
		return "hello"
	}
	//这里调用fn
	fn()
}

func Functional8() {
	//定义一个新的方法，赋值给fn
	fn := func() string {
		return "hello world"
	}() //这个括号，就是匿名方法，需要立即调用
	//如果没下面如果不立刻发起调用，就会报错
	println(fn)
}

// 方法也可以作为返回值
// 下面方法的意思是：返回一个无参数的，返回值为string的方法
func functional7() func() string {
	//return 后面就是这个方法的实现
	return func() string {
		return "hello world"
	}
}
