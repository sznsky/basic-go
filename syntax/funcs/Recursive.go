package main

// 写递归方法，最重要的是要有结束机制,递归使用不当，就可能导致stack overflow
func Recursive(n int) {
	if n > 10 {
		return
	}
	Recursive(n + 1)
}

func A() {
	B()
}

func B() {
	C()
}

func C() {
	A()
}
