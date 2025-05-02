package main

import (
	"dnslog_for_go/internal/log_write"
	"dnslog_for_go/internal/router"
	"dnslog_for_go/internal/web"
)

func main() {
	log_write.InitZapLogger() // 初始化日志
	router.StartServer(web.EmbedFiles)
}
