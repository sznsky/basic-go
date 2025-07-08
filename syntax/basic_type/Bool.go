package main

func Bool() {

	var a bool = true
	var b bool = false
	var c bool = a || b
	println(c)
	var d = a && b
	println(d)
	var e = !a
	println(e)

	println(!(a && b))
	println(!(a || b))

}
