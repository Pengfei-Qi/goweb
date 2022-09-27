package settings

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name          string `mapstructure:"name"`
	Mode          string `mapstructure:"mode"`
	Version       string `mapstructure:"version"`
	Port          int    `mapstructure:"port"`
	*LoggerConfig `mapstructure:"log"`
	*MysqlConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
}

type LoggerConfig struct {
	Level        string `mapstructure:"level"`
	FileName     string `mapstructure:"filename"`
	MaxSize      int    `mapstructure:"max_size"`
	MaxAge       int    `mapstructure:"max_age"`
	MaxBackups   int    `mapstructure:"max_backups"`
	PrintConsole bool   `mapstructure:"print_console"`
	FormatJson   bool   `mapstructure:"format_json"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	//1-测试yaml中读取配置,相对路径
	//viper.SetConfigFile("./conf/config.yaml") // 指定配置文件路径
	//2-测试json中读取配置
	//viper.SetConfigFile("testConfig.json")
	//3-绝对路径读取配置
	//viper.SetConfigFile("E:\\ProjectCode\\go_test\\goweb\\goweb32_bells-of-ireland\\testConfig.json")

	//4. 从路径中读取参数配置,
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.AddConfigPath("./conf") // 在工作目录中查找配置
	//从远程目录中读取配置
	//viper.SetConfigType("yaml")          // 如果配置文件的名称中没有扩展名，则需要配置此项

	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig() is error")
		return err
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		zap.L().Fatal("unable to decode into struct", zap.Error(err))
		return err
	}
	//监听文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(&Conf)
		if err != nil {
			zap.L().Fatal("unable to decode into struct", zap.Error(err))
		}

		fmt.Println("config file changed", in.Name)
	})
	return

}
