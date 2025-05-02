package utils

import (
	"dnslog_for_go/internal/domain/dns_server"
	"dnslog_for_go/internal/log_write"
	"github.com/miekg/dns"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
	"strconv"
	"strings"
	"time"
)

// DNSQueryResult DNS 查询结果结构体
type DNSQueryResult struct {
	Domain  string `json:"domain"`
	IP      string `json:"ip"`
	Address string `json:"address"`
}

var GOLOBAL_PACT = "udp" // 默认协议

// ResolveDNS dns查询
func ResolveDNS(domainName string) DNSQueryResult { // 返回查询结果
	c := &dns.Client{
		Net:     GOLOBAL_PACT,
		Timeout: 10 * time.Second, // 增加超时时间
	}

	message := new(dns.Msg)
	message.SetQuestion(dns.Fqdn(domainName), dns.TypeA) // 查询 A 记录（IPv4）

	r, _, err := c.Exchange(message, getServer()+":53")
	if err != nil {
		log_write.Error("DNS query failed: %v", zap.Error(err))
		return DNSQueryResult{
			Domain:  domainName,
			IP:      "查询失败",
			Address: "无法获取 DNS 服务器",
		}
	}

	// 获取所有 IP 地址
	var ipList []string
	for _, ans := range r.Answer {
		if aRecord, ok := ans.(*dns.A); ok {
			ipList = append(ipList, aRecord.A.String())
		}
	}

	// 真实服务器地址（默认8.8.8.8)
	dnsServer := getServer()
	log_write.Info("正在查询 DNS 服务器", zap.String("server", dnsServer)) // 添加日志

	// 如果没有找到 IPv4 地址，尝试查询 IPv6 地址
	if len(ipList) == 0 {
		message.SetQuestion(dns.Fqdn(domainName), dns.TypeAAAA) // 查询 AAAA 记录（IPv6）
		r, _, err = c.Exchange(message, getServer()+":53")
		if err != nil {
			log_write.Error("DNS query failed for AAAA record: %v", zap.Error(err))
		} else {
			for _, ans := range r.Answer {
				if aaaaRecord, ok := ans.(*dns.AAAA); ok {
					ipList = append(ipList, aaaaRecord.AAAA.String())
				}
			}
		}
	}

	return DNSQueryResult{
		Domain:  domainName,
		IP:      strings.Join(ipList, ", "),
		Address: dnsServer,
	}
}

func getServer() string {
	cfg, err := ini.Load("internal/config/dns_server.ini")
	if err != nil {
		log_write.Error("无法读取配置文件")
		panic("无法读取配置文件")
	}

	current := cfg.Section("DNS").Key("server").String()
	if current == "127.0.0.1" {
		return current
	}

	currentNum, err := strconv.Atoi(current)
	if err != nil {
		log_write.Error("配置值不是有效数字")
		panic("配置值不是有效数字")
	}
	return dns_server.DnsServer(currentNum)
}
