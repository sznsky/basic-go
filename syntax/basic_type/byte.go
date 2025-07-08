package main

import "fmt"

func Byte() {

	var a byte = 'a'
	println(a) //打印的是ascII 码
	//要打出字符串，需要使用格式化打印

	println(fmt.Sprintf("%c", a))

}
