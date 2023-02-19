package etc

import (
	"os"

	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	Name string
	Addr string
	Mode string
}

//初始化配置
func InitConfig(configFile string) error {
	content, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return err
	}

	return nil
}

//获取配置
func GetConfig() *Config {
	return &config
}
