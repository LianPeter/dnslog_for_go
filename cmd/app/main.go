package main

import (
	"dnslog_for_go/internal/router"
	"dnslog_for_go/internal/web"
)

func main() {
	router.StartServer(web.EmbedFiles)
}
