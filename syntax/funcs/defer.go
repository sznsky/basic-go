package main

import "fmt"

//延迟调用defer

func Defer() {

	defer func() {
		println("第一个defer")
	}()

	defer func() {
		println("第二个defer")
	}()
}

func DeferClosure() {
	//调用的时候，先调用的是第二个defer，然后才是第一个defer
	Defer()
}

// 这个是作为闭包传递的
func DeferClosure1() {
	//1.先执行 i:=0
	i := 0
	//3.最后执行defer,所以，最终打印的是1
	defer func() {
		fmt.Println(i)
	}()
	//2.再执行i=1
	i = 1
}

// 注意：下面这种写法，是不一样的【这个是作为参数传递的】
func DeferClosureV1() {
	i := 0
	//这里打印的是0
	defer func(val int) {
		println(val)
	}(i)
	i = 1
}

// defer增加返回值: 这种返回还是0，不会修改生效
func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

// 下面这种修改的会生效：因为这个是带名字的返回值
func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

// 这个会打印出10个10，因为defer是延迟执行的，可以理解为等循环走完了，然后执行的10次defer,
func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

// 这个执行的结果是：9,8,7,...,0
func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

// 和上面是一样的，这个执行的结果是：9,8,7,...,0
func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}
