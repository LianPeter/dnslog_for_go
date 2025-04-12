package main

import (
	"dnslog_for_go/router"
	"embed"
)

//go:embed static/* templates/*
var content embed.FS

func main() {
	router.StartServer(content)
}
