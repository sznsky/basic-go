package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	//通过make创建结构体对象
	personMap := make(map[string]Person)

	personMap["alice"] = Person{Name: "Alice", Age: 21, City: "Wuhan"}
	personMap["lilei"] = Person{Name: "Lilei", Age: 18, City: "Beijing"}
	personMap["wangwu"] = Person{Name: "Wangwu", Age: 22, City: "shanghai"}

	fmt.Println("Entity map:", personMap)

	//访问单个元素
	if p, exists := personMap["alice"]; exists {
		fmt.Printf("Found alice: Name=%s, Age=%d, City=%s\n", p.Name, p.Age, p.City)
	}

	//遍历Map,键值对
	for k, v := range personMap {
		fmt.Printf("key: %s, Name=%s, Age=%d, City=%s\n", k, v.Name, v.Age, v.City)
	}
	//遍历map,只遍历key
	for k := range personMap {
		fmt.Printf("只遍历key: %s\n", k)
	}
	//遍历map,只遍历value
	for _, v := range personMap {
		fmt.Printf("只遍历value: %s\n", v.City)
	}

	//修改Map中某个值，这样不生效
	//personMap["lilei"].Age = 10
	//这样修改才会生效
	personMap["lilei"] = Person{Name: "Lilei", Age: 11, City: "Beijing"}
	for k, v := range personMap {
		fmt.Printf("修改以后的，key: %s, Name=%s, Age=%d, City=%s\n", k, v.Name, v.Age, v.City)
	}

	//删除键值对
	delete(personMap, "wangwu")
	for k, v := range personMap {
		fmt.Printf("删除以后的，key: %s, Name=%s, Age=%d, City=%s\n", k, v.Name, v.Age, v.City)
	}

	//获取map的长度
	length := len(personMap)
	fmt.Printf("map的长度：%d\n", length)

	//检查key是否存在
	value, exists := personMap["alice"]
	if exists {
		fmt.Printf("key exists:", value)
	} else {
		fmt.Printf("key not exists:")
	}

	//
}
