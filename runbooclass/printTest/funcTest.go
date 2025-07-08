package main

// Add 函数
func Add(a, b int) int {
	return a + b
}

// Person 结构体
type Person struct {
	name string
	age  int
}

// Name 方法
func (p Person) Name() string {
	return p.name
}

// Age 方法
func (p Person) Age() int {
	return p.age
}

// GetName 同一个包内，函数也是可以访问的
func GetName(p Person) string {
	return p.name
}

//
