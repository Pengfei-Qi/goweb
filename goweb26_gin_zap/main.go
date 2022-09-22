package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goweb26_gin_zap/config"
	"goweb26_gin_zap/logger"
	"net/http"
	"os"
)

func main() {
	//取不到关键数据-------待解决
	//load config from config.json
	if len(os.Args) < 1 {
		return
	}
	fmt.Printf("打印变量:  %s \n", os.Args)
	if err := config.Init(os.Args[1]); err != nil {
		panic(err)
	}

	// init logger
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v \n", err)
		return
	}

	gin.SetMode(config.Conf.Mode)

	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/hello", func(c *gin.Context) {
		var (
			name = "老六"
			age  = 20
		)
		zap.L().Debug("this is hello func ", zap.String("user", name), zap.Int("age", age))

		c.JSON(http.StatusOK, gin.H{"msg": "hello gin zap framework"})
	})

	addr := fmt.Sprintf(":%v", config.Conf.Port)
	r.Run(addr)
}
