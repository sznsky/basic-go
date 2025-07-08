package main

import "math"

func main() {

	var a int = 456
	var b int = 123
	println(a + b)
	println(a - b)
	println(a * b)
	//注意：做除法的时候，要判断一下是除数
	if b != 0 {
		println(a / b)
		println(a % b)
	}

	//go中不同类型的不能加减乘除,要转行成相同的类型
	var c float64 = 1
	//println(a + c)  //这样写不行的
	println(a + int(c))
	println(float64(a) + c)

	//还有int,int8,int16,int32,int64之间也是不同相互进行运算的
	//var d int32 = 12
	//println(a+d) //这样也是不行的

	//总结：1.go语言中不存在隐式类型转换，不同的类型之间不能进行运算的，这和java是不同的，
	//java中，int和float,double之间是可以进行运算的

	//go的极大值和极小值
	println("float64的最大值：", math.MaxFloat64)
	println("float64的最小正整数：", math.SmallestNonzeroFloat64)

	println("float32的最大值：", math.MaxFloat32)
	println("float32的最小正整数：", math.SmallestNonzeroFloat32)

	//String()

	//Byte()

	Bool()

}
