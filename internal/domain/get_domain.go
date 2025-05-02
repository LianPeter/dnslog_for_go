package domain

import (
	"dnslog_for_go/pkg/utils"
	"github.com/gin-gonic/gin"

	"net/http"
)

var GlobalDomainNameForGetDomain string

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查域名是否合法
	if !utils.StandardizeDomain(domain.DomainName) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "域名不合法，请重新输入"})
		return
	}

	// dns查询
	dnsResult := utils.ResolveDNS(domain.DomainName)

	// 返回查询结果
	if dnsResult.IP == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "没有找到相关的 DNS 记录"})
	} else {
		c.JSON(http.StatusOK, dnsResult)
	}
}

// RandomDomain 随机生成域名
func RandomDomain(c *gin.Context) {
	domainName := GeneratingDomain()
	// GlobalDomainNameForGetDomain = domainName
	// 返回生成的域名
	c.JSON(http.StatusOK, gin.H{"domain": domainName})
}
