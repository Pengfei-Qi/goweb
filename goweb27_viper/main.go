package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"

	"github.com/spf13/pflag"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

/**
参考: https://www.liwenzhou.com/posts/Go/viper_tutorial/

一些场景还没有测试完成
*/

func main() {
	//设置默认值
	//setDefaultValue()

	//测试文件写入
	//testSafeWriteFile()

	//配置文件监听
	//testWatchFile()

	//从io.Reader读取配置
	testReadConfig()

	//数据覆盖
	//testSettingOverride()

	//注册和使用别名
	//testRegisterAlias()

	//使用环境变量, 测试失败, 还是不会
	//testEnvVariable()

	// 标准库Flag
	flag.Int("flagname", 1234, "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	//解析flag
	pflag.Parse()
	//viper绑定pflag
	viper.BindPFlags(pflag.CommandLine)

	fmt.Printf("variable filename is %d \n", viper.GetInt("flagname"))

	//启动服务
	//startServer()
}

func setDefaultValue() {
	viper.SetConfigFile("./config.yaml")  // 指定配置文件路径
	viper.SetConfigName("config")         // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")           // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在的路径
	viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".")              // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()           // 查找并读取配置文件
	if err != nil {                       // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func testEnvVariable() {
	//SetEnvPrefix("spf") // 将自动转为大写
	//BindEnv("id")
	//
	//os.Setenv("SPF_ID", "13") // 通常是在应用程序之外完成的
	//
	//id := Get("id") // 13
}

func testRegisterAlias() {
	viper.RegisterAlias("loud", "Verbose") // 注册别名（此处loud和Verbose建立了别名）

	viper.Set("verbose", true) // 结果与下一行相同
	viper.Set("loud", true)    // 结果与前一行相同

	fmt.Printf("print var loud :%t\n", viper.GetBool("loud"))
	fmt.Printf("print var verbose :%t\n", viper.GetBool("verbose"))
}

func testSettingOverride() {
	viper.Set("beard", false)
	fmt.Printf("[Setting Overrides]print var beard :%t\n", viper.GetBool("beard"))
}

// testWatchFile 测试监听文件
func testWatchFile() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
}

func startServer() {
	r := gin.Default()
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})

	r.Run(":8091")
}

func testSafeWriteFile() {
	//1. 测试写入文件
	viper.WriteConfig() // 将当前配置写入“viper.AddConfigPath()”和“viper.SetConfigName”设置的预定义路径
	viper.SafeWriteConfig()
	viper.WriteConfigAs("./testViper")
	//viper.SafeWriteConfigAs("./testViper") // 因为该配置文件写入过，所以会报错
	viper.SafeWriteConfigAs("./otherViper")
}

func testReadConfig() {
	viper.SetConfigType("yaml") // 或者 viper.SetConfigType("YAML")

	// 任何需要将此配置添加到程序中的方法。
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	fmt.Printf("print var eyes :%s\n", viper.Get("eyes"))
	fmt.Printf("print var beard :%t\n", viper.GetBool("beard"))
}
