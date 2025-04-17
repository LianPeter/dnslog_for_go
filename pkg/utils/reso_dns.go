package utils

import (
	"dnslog_for_go/internal/log_write"
	"github.com/miekg/dns"
	"go.uber.org/zap"
	"strings"
	"time"
)

// DNSQueryResult DNS 查询结果结构体
type DNSQueryResult struct {
	Domain  string `json:"domain"`
	IP      string `json:"ip"`
	Address string `json:"address"`
}

// ResolveDNS dns查询
func ResolveDNS(domainName string) DNSQueryResult { // 返回结构体即dns查询结果
	c := &dns.Client{
		Net:     "udp",
		Timeout: 5 * time.Second, // 设置超时时间
	}

	message := new(dns.Msg)
	message.SetQuestion(dns.Fqdn(domainName), dns.TypeA)

	// 8.8.8.8:53可以更换
	r, _, err := c.Exchange(message, "8.8.8.8:53")
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

	// 真实服务器地址（从 DNS 响应中尝试获取）
	dnsServer := "8.8.8.8" // 实际用于查询的地址

	return DNSQueryResult{
		Domain:  domainName,
		IP:      strings.Join(ipList, ", "),
		Address: dnsServer,
	}
}
