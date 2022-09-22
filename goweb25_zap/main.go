package main

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

// InitLogger 第一种- 采用默认的logger
//func InitLogger() {
//	logger, _ = zap.NewProduction()
//	sugarLogger = logger.Sugar()
//}

// InitLogger 第二种, 采用自定义全量日志输出
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	//增加调用者信息
	//logger := zap.New(core, zap.AddCaller())
	//简化调用链
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	//sugarLogger = logger.Sugar()
}

// InitLogger 第三种, 全量日志输出到文件, 同时提取error日志到 err.log中
//func InitLogger() {
//	encoder := getEncoder()
//	//全量日志到test.log中
//	logF, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
//	c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), zapcore.DebugLevel)
//	//test.err.log记录ERROR级别的日志
//	errF, _ := os.OpenFile("./test.err.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
//	c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zapcore.ErrorLevel)
//
//	//使用NewTee将c1和c2合并到core
//	core := zapcore.NewTee(c1, c2)
//	//简化调用链
//	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
//	sugarLogger = logger.Sugar()
//}

func getEncoder() zapcore.Encoder {
	//第一种方式
	//encoderConfiger := zap.NewProductionEncoderConfig()
	//encoderConfiger.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfiger.EncodeLevel = zapcore.CapitalLevelEncoder
	//第二种方式
	encoderConfiger := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zapcore.NewConsoleEncoder(encoderConfiger)
}

func getLogWriter() zapcore.WriteSyncer {
	//第一种, 仅输出到文件
	//file, _ := os.Create("./test.log") // 每次新建文件并写入日志
	//file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	//return zapcore.AddSync(file)

	//第二种, 在控制台和日志文件同步输出
	//ws := io.MultiWriter(file, os.Stdout)
	//return zapcore.AddSync(ws)

	//第三种, 日志输出到文件, 同时按照大小切割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,     //单位:M
		MaxBackups: 5,     //保留旧文件的最大个数
		MaxAge:     30,    //单位: 天
		Compress:   false, //是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

//func main() {
//	InitLogger()
//	defer sugarLogger.Sync()
//
//	//测试写入日志
//	//for i := 0; i < 10000; i++ {
//	//	sugarLogger.Infof("[warning 测试日志文件切割] \n 君不见，黄河之水天上来，奔流到海不复回。\n君不见，高堂明镜悲白发，朝如青丝暮成雪。\n人生得意须尽欢，莫使金樽空对月。\n天生我材必有用，千金散尽还复来。\n烹羊宰牛且为乐，会须一饮三百杯")
//	//}
//
//	//simpleHttpGet("https://www.baidu.com/")
//	//simpleHttpGet("www.baidu.com")
//
//
//}

func main() {
	//	日志初始化
	InitLogger()
	defer logger.Sync()

	r := gin.New()
	r.Use(GinLogger(), GinRecovery(true))
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "法外狂徒张三....王老五"})
	})
	r.Run()
}

//在 gin 中实现zap日志库

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
