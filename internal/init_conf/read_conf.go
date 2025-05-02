package init_conf

import (
	"dnslog_for_go/internal/config"
	"dnslog_for_go/internal/log_write"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
)

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
	config.GlobalDomainNameForReadConfig = readDomain

	log_write.Info("读取配置文件成功", zap.String("domain", readDomain))

}
