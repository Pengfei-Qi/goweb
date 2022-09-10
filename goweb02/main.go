package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()

	engine.GET("/helloGo", sayHello)

	engine.Run(":9000")
}

func sayHello(context *gin.Context) {
	context.JSON(200, gin.H{
		"massage": "hello goland! Let's go!",
	})
}
