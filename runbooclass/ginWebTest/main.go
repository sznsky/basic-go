package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个gin路由器
	r := gin.Default()

	// 定义一个get请求路径，并返回数据
	r.GET("/hello", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"message": "hello world"})
		c.String(http.StatusOK, "hello world")
	})

	// 启动服务，监听8080端口, 如果有错误，打印
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
