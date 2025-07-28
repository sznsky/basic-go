package main

import (
	"github.com/gin-gonic/gin"
)

// Index 路由具体实现的方法
func Index(c *gin.Context) {
	c.String(200, "hello gin")
}

func main() {
	// 创建一个默认的路由
	router := gin.Default()
	router.GET("/index", Index)
	//启动方式1
	router.Run(":8080")
	// 启动方式2,router.Run 就是对http.ListenAndServe的封装
	//http.ListenAndServe(":8080", router)
}
