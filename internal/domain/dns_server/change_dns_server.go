package dns_server

import (
	"dnslog_for_go/internal/log_write"
	"gopkg.in/ini.v1"
	"os"
	"strconv"
)

// Change 修改 DNS 服务器
// num 设置的dns服务器编号
func Change(num byte) {
	dir, _ := os.Getwd()
	log_write.Info("当前工作目录:" + dir)

	// 读取配置文件
	cfg, err := ini.Load("internal/config/dns_server.ini")
	if err != nil {
		log_write.Error("无法读取配置文件")
		panic("Unable to read configuration file")
	} else {
		log_write.Info("读取配置文件成功")
	}

	// 获取当前 DNS 设置
	current := cfg.Section("DNS").Key("server").String()
	currentNum, err := strconv.Atoi(current)
	if err != nil {
		log_write.Error("配置值不是有效数字")
		panic("Configuration values are not valid numbers")
	}

	// 如果当前 DNS 已经是目标 DNS，则无需修改
	if int(num) == currentNum {
		log_write.Info("DNS 设置已是当前值，无需修改")
		return
	}

	// 设置新的 DNS
	setDnsErr := setDNS(strconv.Itoa(int(num)))
	if setDnsErr != nil {
		return
	}

	// 更新配置文件
	cfg.Section("DNS").Key("server").SetValue(strconv.Itoa(int(num)))
	err = cfg.SaveTo("internal/config/dns_server.ini")
	if err != nil {
		log_write.Error("保存配置失败")
		panic("Failed to save configuration")
	} else {
		log_write.Info("保存配置成功")
	}

	// 日志记录修改后的 DNS
	log_write.Info("DNS 设置已更新为: " + strconv.Itoa(int(num)))
}

// setDNS 通过系统命令设置 DNS
func setDNS(value string) error {
	cfg, err := ini.Load("internal/config/dns_server.ini")
	if err != nil {
		return err
	}

	cfg.Section("DNS").Key("server").SetValue(value)
	return cfg.SaveTo("internal/config/dns_server.ini")
}
