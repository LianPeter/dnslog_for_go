package utils

import (
	"dnslog_for_go/internal/log_write"
	"gopkg.in/ini.v1"
)

func SelectPact(s string) string {
	cfg, err := ini.Load("internal/config/dns_server.ini")
	if err != nil {
		log_write.Error("无法读取配置文件")
		panic("Unable to read configuration file")
	} else {
		log_write.Info("读取配置文件成功")
	}

	return cfg.Section("PACT").Key(s).String()
}
