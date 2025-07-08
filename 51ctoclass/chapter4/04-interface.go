package main

import "fmt"

// 定义动物接口，两个动作
type Animal interface {
	sleeping()
	Eating()
}

type Dog struct {
	color string
}
type Cat struct {
	color string
}

// dog支持接口
func (d Dog) sleeping() {
	fmt.Printf("%s's dog is sleeping\n", d.color)
}
func (d Dog) Eating() {
	fmt.Printf("%s's dog is Eating\n", d.color)
}

// cat支持接口
func (c Cat) sleeping() {
	fmt.Printf("%s's cat is sleeping\n", c.color)
}
func (c Cat) Eating() {
	fmt.Printf("%s's cat is Eating\n", c.color)
}

// 工厂模式；流水线的作业，可以生产猫，也可以生产狗，返回值怎么办
func factory(color, animal_type string) Animal {
	switch animal_type {
	case "dog":
		return &Dog{color: color}
	case "cat":
		return &Cat{color: color}
	default:
		return nil
	}
}
func main() {
	// 不是工厂模式
	dog := Dog{"white"}
	dog.sleeping()
	dog.Eating()

	cat := Cat{"black"}
	cat.sleeping()
	cat.Eating()

	// 使用工厂模式，一样可以实现的上面的效果
	c1 := factory("yellow", "cat")
	c1.sleeping()
	c1.Eating()

	c2 := factory("green", "dog")
	c2.sleeping()
	c2.Eating()

}
