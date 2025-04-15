package domain

import (
	"dnslog_for_go/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/miekg/dns"
	"log"
	"net/http"
	"strings"
	"time"
)

var GlobalDomainNameForGetDomain string

// DNSQueryResult DNS 查询结果结构体
type DNSQueryResult struct {
	Domain  string `json:"domain"`
	IP      string `json:"ip"`
	Address string `json:"address"`
}

// ShowForm 展示表单
func ShowForm(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// SubmitDomain 提交域名并查询
func SubmitDomain(c *gin.Context) {
	var domain struct {
		DomainName string `json:"domain_name"` // 接收域名
	}

	// 解析JSON数据
	if err := c.ShouldBindJSON(&domain); err != nil {
		// 解析出错，返回400 错误
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 使用校验函数检查域名是否合法
	if !utils.StandardizeDomain(domain.DomainName) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "域名不合法，请重新输入"})
		return
	}

	// dns查询
	dnsResult := resolveDNS(domain.DomainName)

	// 返回查询结果
	if dnsResult.IP == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "没有找到相关的 DNS 记录"})
	} else {
		c.JSON(http.StatusOK, dnsResult)
	}
}

// resolveDNS dns查询
func resolveDNS(domainName string) DNSQueryResult { // 返回结构体即dns查询结果
	c := &dns.Client{
		Net:     "udp",
		Timeout: 5 * time.Second, // 设置超时时间
	}

	message := new(dns.Msg)
	message.SetQuestion(dns.Fqdn(domainName), dns.TypeA)

	// 8.8.8.8:53可以更换
	r, _, err := c.Exchange(message, "8.8.8.8:53")
	if err != nil {
		log.Println("DNS query failed:", err)
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
