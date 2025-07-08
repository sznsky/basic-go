package main

// 包变量
var Global = "全局变量"
var internal = "包含内变量，私有变量"

//总结：全局变量首字母大写才能被其他的函数访问的到，比如上面Global可以被其他的函数访问，但是internal不能

// 包变量块:和上面全局变量一样，首字母大写的变量才能被范围的到，否则访问不到
var (
	Fisrt  string = "1" //可以被其他函数访问
	second int    = 2   //不能被其他函数访问
)

// 局部变量的声明
func main() {

	//默认是int,可以省略
	var a int = 123
	println(a)

	var b = 456
	println(b)

	//如果声明的不是默认类型，是不能省略的
	var c uint = 456
	println(c)

	//注意:go中不同类型直接，是不能进行运算，比较的，比如下面的，都不行
	//println( a+c)
	//println( a == c )

	//定义变量块
	var (
		d string = "aaa"
		e int    = 123
	)
	println(d, e)

	//冒等声明变量，只能是局部变量，同时go会进行类型推断，
	//注意:数字只能是int或者float64类型,其他类型只能使用var
	f := 1
	println(f)
	g := 1.2
	println(g)
	h := "你好"
	println(h)

	//同一个变量，只能声明一次（在一个作用域内），但是可以再次赋值
	var i = 1 //声明并赋值
	i = 123   //再次赋值
	println(i)

	var j int //仅仅声明，但是这种写法，类型不能被推断，要直接写出类型
	j = 1     //赋值
	println(j)

	//局部变量和全局变量可以声明称同一个名称，但是会被局部变量覆盖掉
	var Fisrt = "123" //注意：绿色，表明覆盖掉了，ide提示，一般不要这么写
	println(Fisrt)

	//

}

//注意：1.go是支持类型推断的，整数默认是int类型，小数默认是float64类型
//2.局部变量没有使用，会编译不通过
