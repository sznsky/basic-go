package main

// go写个简单web
import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	// 判断路径，只处理hello的请求
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	// 只处理GET请求
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.Must Get method", http.StatusMethodNotAllowed)
		return
	}

	// 将返回数据绑定到ResponseWriter上，并返回
	_, err := fmt.Fprintf(w, "Hello World,I'm golang")
	if err != nil {
		return
	}
}

func main() {
	// 注册处理函数到特定路径
	http.HandleFunc("/hello", helloHandler)

	// 启动日志
	fmt.Println("Starting server at port 8080")

	//启动web服务器，监听端口 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
