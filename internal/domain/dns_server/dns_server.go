package dns_server

func DnsServer(num int) string {
	server := []string{
		"8.8.8.8",   // Google Public DNS
		"223.5.5.5", // 阿里公共 DNS
		"127.0.0.1", // 本地 DNS
	}
	return server[num]
}
