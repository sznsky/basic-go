package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _string(c *gin.Context) {
	c.String(http.StatusOK, "你好啊")
}

func _json(c *gin.Context) {
	type UserInfo struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
		Password string `json:"-"` // 不要转换为json
	}
	user := UserInfo{"lilei", 10, "123456"}
	c.JSON(http.StatusOK, user)
}

func main() {
	router := gin.Default()
	router.GET("/", _string)
	router.GET("/json", _json)
	router.Run(":80")
}
