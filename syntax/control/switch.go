package main

// 在go中，switch每个case后面是不需要break的
func Switch(status int) {
	switch status {
	case 0:
		println("初始化")
	case 1:
		println("运行中")
	default:
		println("未知运行状态")
	}
}

// 还可以这样写, switch后面不带参数，就是Bool类型的
func SwitchBool(age int) {
	switch {
	case age > 18:
		println("成年人")
	case age > 12:
		println("青少年")
	default:
		println("儿童")
	}
}
