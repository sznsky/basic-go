package main

func IfOnly(age int) {
	if age >= 18 {
		println("成年了...")
	}
}

func IfElse(age int) {
	if age >= 18 {
		println("成年了...")
	} else if age > 10 {
		println("青少年...")
	} else {
		println("儿童...")
	}
}

func IfElseReturn(age int) string {
	if age > 18 {
		return "成年了"
	} else if age > 10 {
		return "青少年"
	} else {
		return "儿童"
	}
}

// 这个写法是在其他语言没有的写法
func IfNewVariable(start int, end int) string {
	if distance := end - start; distance > 100 {
		return "太远了"
	} else if distance > 60 {
		return "有点远"
	} else {
		return "还挺好"
	}
}
