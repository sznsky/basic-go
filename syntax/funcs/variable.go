package main

// 参数name,和不定参数aliases
func YourName(name string, aliases ...string) {

}

// 不定参数，可以传入一个，两个，或者多个，也可以不传
func CallYourName() {
	YourName("大明")
	YourName("大明", "小明")
	YourName("大明", "小明", "李雷")

	//定义一个切片，将切片传入不定参数中
	aliases := []string{"小明", "大明"}
	YourName("大明", aliases...)
}
