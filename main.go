package main

import (
	"github.com/gin-gonic/gin"
)

// 实现对  "/hello/:name" 的处理器
// 方法签名 HandlerFunc type HandlerFunc func(*Context)
func handlerName(c *gin.Context) {
	name := c.Params.ByName("name")
	c.String(200, name)
}

// 入口函数
func main() {
	r := gin.Default()
	//创建路由对应处理器
	r.GET("/hello/:name", handlerName)
	//启动服务
	r.Run("localhost:8080")
}
