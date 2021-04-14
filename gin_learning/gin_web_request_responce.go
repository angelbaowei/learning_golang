package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	username := c.DefaultQuery("username", "miaozhibin")
	address := c.Query("address")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!!!",
		"username": username,
		"address": address,

	})
	// http://localhost:8888/hello?username=miaozhibin2&address=jiyuan
}

func main() {
	r := gin.Default()  // 返回默认的路由引擎

	r.GET("/hello", sayHello)  // 请求/hello 执行sayHello函数
	r.GET("/test", func(c *gin.Context) {  // http重定向
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})
	r.GET("/test2", func(c *gin.Context) {  // 路由重定向
		// 指定重定向的URL
		c.Request.URL.Path = "/test"
		r.HandleContext(c)
	})

	r.NoRoute(func(c *gin.Context) {  // 为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码
		c.JSON(http.StatusNotFound, gin.H{
			"message": "404",
		})
	})
	//r.Any("/test", func(c *gin.Context) {...}) // 还有一个可以匹配所有请求方法的Any方法如下

	//r.POST()  增
	//r.DELETE() 删
	//r.PUT() 改
	//r.GET() 查

	// 启动服务
	r.Run(":8888")
}
