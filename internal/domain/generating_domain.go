package domain

import (
	"dnslog_for_go/internal/log_write"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strings"
)

// GeneratingDomain 基于uuid生成域名
func GeneratingDomain() string {
	commonTLDs := []string{
		".com", ".net", ".org", ".cn", ".io", ".edu", ".gov", ".co", ".xyz",
	}
	id := uuid.New().String()

	cleaned := strings.ReplaceAll(id, "-", "")

	if len(cleaned) < 10 {
		return fmt.Sprintf("UUID 过短，不足10字符: %s", cleaned)
	}
	shortDomain := cleaned[:10]

	log_write.Info("生成的短域名为: " + shortDomain)

	if len(shortDomain) > 10 {
		return fmt.Sprintf("域名长度超过限制: %s", shortDomain)
	}

	i := rand.Intn(9)

	tld := commonTLDs[i]
	domain := fmt.Sprintf("%s%s", shortDomain, tld)

	log_write.Info("完整的域名为: " + domain)
	return domain
}
