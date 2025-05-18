package dns_server

import (
	"dnslog_for_go/pkg/log"
	"gopkg.in/ini.v1"
	"os"
	"strconv"
)

// ChangeServer 修改 DNS 服务器
// num 设置的dns服务器编号
func ChangeServer(num byte) {
	dir, _ := os.Getwd()
	log.Info("当前工作目录:" + dir)

	cfg, err := ini.Load("internal/config/dns_server.ini")
	if err != nil {
		log.Error("无法读取配置文件")
		panic("Unable to read configuration file")
	} else {
		log.Info("读取配置文件成功")
	}

	current := cfg.Section("DNS").Key("server").String()
	currentNum, err := strconv.Atoi(current)
	if err != nil {
		log.Error("配置值不是有效数字")
		panic("Configuration values are not valid numbers")
	}

	if int(num) == currentNum {
		log.Info("DNS 设置已是当前值，无需修改")
		return
	}

	setDnsErr := setDNS(strconv.Itoa(int(num)))
	if setDnsErr != nil {
		return
	}

	// 更新配置文件
	cfg.Section("DNS").Key("server").SetValue(strconv.Itoa(int(num)))
	err = cfg.SaveTo("internal/config/dns_server.ini")
	if err != nil {
		log.Error("保存配置失败")
		panic("Failed to save configuration")
	} else {
		log.Info("保存配置成功")
	}

	log.Info("DNS 设置已更新为: " + strconv.Itoa(int(num)))
}

// setDNS 设置 DNS 服务器
func setDNS(value string) error {
	cfg, err := ini.Load("internal/config/dns_server.ini")
	if err != nil {
		return err
	}

	cfg.Section("DNS").Key("server").SetValue(value)
	return cfg.SaveTo("internal/config/dns_server.ini")
}
