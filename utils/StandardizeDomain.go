package utils

import (
	"log"
	"regexp"
	"strings"
)

// 允许的常见 TLD 后缀（可扩展）
var commonTLDs = []string{
	".com", ".net", ".org", ".cn", ".io", ".edu", ".gov", ".co", ".xyz",
}

// StandardizeDomain 判断域名是否合法
func StandardizeDomain(domain string) bool {
	// 转成小写
	domain = strings.ToLower(domain)

	// 长度限制
	if len(domain) < 3 || len(domain) > 253 {
		log.Println("域名长度必须在 3 到 253 字符之间")
		return false
	}

	// 检查是否有合法的 TLD
	validTLD := false
	for _, tld := range commonTLDs {
		if strings.HasSuffix(domain, tld) {
			validTLD = true
			break
		}
	}
	if !validTLD {
		log.Println("域名必须以常见的顶级域名结尾，例如 .com、.net")
		return false
	}

	// 正则校验整个域名结构是否合法
	match, _ := regexp.MatchString(`^(?i:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)(?:\.(?i:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?))*\.[a-z]{2,}$`, domain)
	if !match {
		log.Println("域名结构不合法。允许多个标签，中间用点分隔，每个标签仅由字母数字和中划线组成，且不能以中划线开头或结尾")
		return false
	}

	log.Println("域名符合规范!")
	return true
}
