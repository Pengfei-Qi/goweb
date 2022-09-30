package main

import (
	"fmt"
	"goweb32_bells-of-ireland/controllers/controller"
	"goweb32_bells-of-ireland/dao/mysql"
	"goweb32_bells-of-ireland/dao/redis"
	"goweb32_bells-of-ireland/logger"
	"goweb32_bells-of-ireland/pkg/snowflake"
	"goweb32_bells-of-ireland/routers"
	"goweb32_bells-of-ireland/settings"

	"go.uber.org/zap"
)

func main() {

	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("setting init failed, err : %v \n", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(settings.Conf.LoggerConfig); err != nil {
		fmt.Printf("logger init failed, err : %v \n", err)
		return
	}
	defer zap.L().Sync()
	//3.初始化MySql连接
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		zap.L().Info("mysql init failed, err : %v \n", zap.Error(err))
		return
	}
	defer mysql.Close()
	//4.初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		zap.L().Info("mysql init failed, err : %v \n", zap.Error(err))
		return
	}
	defer redis.Close()
	//5.初始化雪花算法连接
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		zap.L().Info("snowflake init failed, err : %v \n", zap.Error(err))
		return
	}
	//6.初始化雪花算法
	if err := controller.InitTrans("zh"); err != nil {
		zap.L().Info("snowflake init failed, err : %v \n", zap.Error(err))
		return
	}

	//最后.注册路由
	routers.SetUp()
}
