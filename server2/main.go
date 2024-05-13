package main

import (
	"fmt"
	"net/http"
	"server2/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	// 接收命令行参数
	configFile := pflag.StringP("config", "c", "../config/config1.yaml", "choose config file.") // 指定配置文件
	pflag.Parse()

	// 解析配置文件
	v := viper.New()
	v.SetConfigFile(*configFile)
	err := v.ReadInConfig()
	if err != nil {
		fmt.Errorf("fatal error config file: %s", err)
		return
	}
	var Config *config.Config

	if err := v.Unmarshal(&Config); err != nil {
		fmt.Errorf("unmarshal config file error: %s", err)
		return
	}

	port := Config.System.TcpPort // 获取 system 下的 tcp_port 配置

	// 初始化 Gin
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "提供v2版本服务",
		})
	})

	// 使用配置文件中的端口启动服务
	r.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
