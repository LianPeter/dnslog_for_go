package dns_server

import (
	"dnslog_for_go/internal/log_write"
	"gopkg.in/ini.v1"
)

func DefaultConfig() {
	cfg, err := ini.Load("internal/config/dns_server.ini")
	if err != nil {
		panic("无法读取配置文件")
	}
	cfg.Section("DNS").Key("server").SetValue("0")

	err = cfg.SaveTo("internal/config/dns_server.ini")
	if err != nil {
		panic("默认配置恢复失败")
	} else {
		log_write.Info("默认配置恢复成功")
	}
}
