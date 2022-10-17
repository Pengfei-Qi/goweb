package routers

import (
	"context"
	"fmt"
	"goweb32_bells-of-ireland/controllers/controller"
	"goweb32_bells-of-ireland/logger"
	"goweb32_bells-of-ireland/middlewares"
	"goweb32_bells-of-ireland/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	router.Use(logger.GinLogger(), logger.GinRecovery(true))

	//根路径
	router.GET("/", func(c *gin.Context) {
		controller.ResponseSuccess(c, gin.H{
			"version":     settings.Conf.Version,
			"projectName": settings.Conf.Name,
		})
	})
	//无路由
	router.NoRoute(func(c *gin.Context) {
		controller.ResponseError(c, controller.CodeServerNotFound)
	})

	v1 := router.Group("/api/v1")

	//用户注册
	v1.POST("/signup", controller.SignUpHandler)

	//用户登陆
	v1.POST("/login", controller.LoginHandler)

	//刷新Token
	v1.GET("/refreshToken", controller.RefreshToken)

	v1.Use(middlewares.JWTAuthMiddleware())

	{
		//ping
		v1.GET("/ping", func(c *gin.Context) {
			controller.ResponseSuccess(c, "pong")
		})

		//获取community
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		//发帖
		v1.POST("/post", controller.CreateArticleHandler)
		v1.GET("/post/:id", controller.QueryArticleDetailHandler)
	}

	//启动服务或者延迟5秒关机
	startOrDelayStopServer(router)

	return router
}

func startOrDelayStopServer(router *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: router,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
