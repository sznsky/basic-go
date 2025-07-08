package main

func main() {
	//方法调用，使用两个参数接受返回值
	//name, age := Fun10()
	//println(name, age)

	//方法调用，直接收一个返回值,另外一个返回值使用下划线代替
	//name1, _ := Fun10()
	//println(name1)

	//方法的调用，:=的前提是，就是左边还有一个新的变量
	//name1, _ := Fun10() //这样写就是错误的
	//name1, name2 := Fun10()
	//println(name1)
	//println(name2)

	//Fun6("hello", "world")

	//Recursive(10)

	//调用函数式方法
	//UseFunctional4()
	//调用原来方法
	//Functional4()

	//匿名方法的调用,这个functional8内部有个匿名方法
	//Functional8()

	//调用closure这个方法，其实它返回的是一个方法
	//fn := Closure("李雷")
	//调用你这个方法fn,打印李雷
	//println(fn())

	//DeferClosure()

	//getAge := Closure1()
	//println(getAge())
	//println(getAge())
	//println(getAge())

	//getAge = Closure1()
	//println(getAge())
	//println(getAge())
	//println(getAge())

	//DeferClosure1()

	//println(DeferReturn())
	//println(DeferReturnV1())

	//DeferClosureLoopV1()

	//DeferClosureLoopV2()

	DeferClosureLoopV3()

}
