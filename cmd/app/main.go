package main

import (
	"dnslog_for_go/internal/domain/dns_server"
	"dnslog_for_go/internal/init_conf"
	"dnslog_for_go/internal/log_write"
	"dnslog_for_go/internal/router"
	"dnslog_for_go/internal/web"
)

func main() {
	log_write.InitZapLogger() // 初始化日志
	router.StartServer(web.EmbedFiles)
	init_conf.IsExist()
	dns_server.DefaultConfig() // 恢复默认配置
}
