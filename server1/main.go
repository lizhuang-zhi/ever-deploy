package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 接收命令行参数
	configPath := flag.String("config", ".", "path to config file")
	flag.Parse()

	// 初始化 viper
	v := viper.New()

	// 设置配置文件的名称
	v.SetConfigName("config1") // 注意这里不带文件扩展名
	// 设置读取的文件类型
	v.SetConfigType("yaml")

	// 使用命令行参数指定的路径
	v.AddConfigPath(*configPath) // 配置文件的路径

	// 读取配置数据
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 获取嵌套配置信息
	port := v.GetString("system.tcp_port") // 获取 system 下的 tcp_port 配置

	// 初始化 Gin
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "提供v1版本服务",
		})
	})

	// 使用配置文件中的端口启动服务
	r.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
