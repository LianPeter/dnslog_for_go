package example

import (
	"net"
	"strings"
)

// StandardizeDomain 校验域名合法性
func StandardizeDomain(domain string) bool {
	return strings.Contains(domain, ".") && !strings.ContainsAny(domain, " <>/:")
}

type DNSResult struct {
	Domain string `json:"domain"`
	IP     string `json:"ip"`
}

// ResolveDNS 模拟 DNS 查询，返回 IP（真实项目中可以更复杂）
func ResolveDNS(domain string) DNSResult {
	ips, err := net.LookupIP(domain)
	if err != nil || len(ips) == 0 {
		return DNSResult{Domain: domain, IP: ""}
	}
	return DNSResult{Domain: domain, IP: ips[0].String()}
}
