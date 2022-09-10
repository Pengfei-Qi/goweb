package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func homeFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "home",
	})
}

func m1(c *gin.Context) {
	fmt.Println("enter m1 .......in")
	c.Set("name", "张三吃席啦~~~~~~~~")
	now := time.Now()
	c.Next()
	cast := time.Since(now)
	fmt.Printf("共耗时:%s\n", cast)
	fmt.Println("enter m1 .......out")
}

func m2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("enter m2 .......in")
		//处理下一个流程
		c.Next()

		fmt.Println("enter m2 .......out")
	}
}

func middleAuth(t bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		name, ok := c.Get("name")
		if !ok {
			name = "张三不在, 李思思来吃席啦"
		}
		if t {
			// 执行业务逻辑
			//c.Next()
			// 不调用该请求的剩余处理程序
			//c.Abort()

			fmt.Printf("正常处理权限校验. 收到的数据为: %s.......完成\n", name)

		} else {
			// 不调用该请求的剩余处理程序
			c.Abort()
			name = "张三不在, 老五来吃席啦"
			fmt.Printf("异常处理权限校验. 收到的数据为: %s.......退出\n", name)
			c.JSON(http.StatusOK, gin.H{"errorMsg": "校验失败, 请求终止"})
			return
		}
	}
}

func main() {
	r := gin.Default()
	//全局注册中间件, 分别采用函数, 闭包两种方式
	r.Use(m1, m2())
	r.GET("/home", homeFunc)
	//模拟权限校验过程, false
	r.GET("/index", middleAuth(false), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"url": c.Request.URL.Path,
		})
	})
	r.GET("/shop", middleAuth(true), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"url": c.Request.URL.Path,
		})
	})

	//为路由组添加中间件
	//写法1
	companyGroup := r.Group("/company", middleAuth(true))

	{
		companyGroup.GET("/unicom", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"url": c.Request.URL.Path,
			})
		})
		companyGroup.GET("/cmcc", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"url": c.Request.URL.Path,
			})
		})
	}

	//写法2
	regionGroup := r.Group("/region")
	regionGroup.Use(middleAuth(false))
	{
		regionGroup.GET("/minhang", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"url": c.Request.URL.Path,
			})
		})
		regionGroup.GET("/pudong", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"url": c.Request.URL.Path,
			})
		})
	}

	r.Run() //默认8080端口
}
