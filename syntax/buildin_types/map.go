package main

// map
func Map() {

	//直接定义一个map
	m1 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	m1["hello"] = 4

	//使用make方法定义map,可以指定容量，也可以不指定，默认是16
	m2 := make(map[string]int, 4)
	m2["hello"] = 4

}
