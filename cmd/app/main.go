package main

import (
	"dnslog_for_go/internal/router"
	"dnslog_for_go/internal/web"
	"dnslog_for_go/pkg/log"
)

func main() {
	log.InitZapLogger() // 初始化日志
	router.StartServer(web.EmbedFiles)
}
