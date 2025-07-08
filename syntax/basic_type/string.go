package main

import (
	"fmt"
	"unicode/utf8"
)

// String 注释 /*
func String() {

	println("hello,go")
	//``这个符号也可以表示字符串，同时还可以换行，比较长的字符串可以这样写
	println(`呵呵呵呵
       斤斤计较 ""
         急急急急急急`)

	//如何想输出  he said:"HELLO GO!"   这个时候需要转义 ，系统自动转义
	println("he said:\"HELLO GO!\" ")

	//字符串的拼接:go只是支持字符串类型之间的拼接，但是java支持字符串和其他类型的拼接
	println("hello" + ",Go!")

	//比如支持：字符串和数字的拼接，先要把数字转换成字符串,可以使用下面的方法
	var s = fmt.Sprintf("hello %d", 123)
	println(s)

	//len获取字符串的长度(这里默认输出的字节的长度，一个中文默认是2个字节 )
	println(len("abc")) //3
	println(len("你好"))  //6
	//如果要获取中文的长度:进行uft8的转换
	println(utf8.RuneCountInString("你好"))

	//字符串api操作都在strings里面

	//字节的操作都在bytes里面

}
