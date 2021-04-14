package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。
//这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。
// Gin中的中间件必须是一个gin.HandlerFunc类型

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

//
func StatCost2() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子2") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	// 注册一个全局中间件
	r.Use(StatCost())

	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	// 给/test2路由单独注册中间件（可注册多个）
	r.GET("/test2", StatCost2(), func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world2!",
		})
	})

	r.Run(":8888")

	//2021/04/14 17:30:34 小王子
	//2021/04/14 17:30:34 66.589µs
	//2021/04/14 17:30:33 小王子2
	//2021/04/14 17:30:33 49.463µs
	//2021/04/14 17:30:33 62.489µs  有两个 一个单独的test2的中间件 一个全局的test的中间件

	/*
	gin.Default()默认使用了Logger和Recovery中间件，其中：
	Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
	Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
	如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。
	 */
}