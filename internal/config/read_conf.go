package config

import (
	"github.com/spf13/viper"
	"log"
)

// GlobalDomainNameForReadConfig 定义全局变量
var GlobalDomainNameForReadConfig string

// ReadConfig 读取配置文件
func ReadConfig() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config.yaml failed: %v", err)
	}

	readDomain := viper.GetString("domain")
	GlobalDomainNameForReadConfig = readDomain

	log.Println("readDomain:", GlobalDomainNameForReadConfig)

}
