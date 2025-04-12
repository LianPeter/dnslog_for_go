package main

import "embed"

//go:embed templates/* static/*
var EmbedFiles embed.FS
