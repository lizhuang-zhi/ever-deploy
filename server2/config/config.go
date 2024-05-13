package config

// Config 配置文件结构体
type Config struct {
	System struct {
		TcpPort string `mapstructure:"tcp_port"`                // TCP 端口
		TmpDir  string `mapstructure:"tmp_dir" usage:"tmp_dir"` // 临时数据存放目录
	} `mapstructure:"system"`
}
