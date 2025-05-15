package log

import (
	"log"
	"os"
	"path/filepath"
)

func findProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir // 找到了 go.mod，就认为是项目根
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			log.Fatal("未找到项目根目录（缺少 go.mod）")
		}
		dir = parent
	}
}

func InitLog() {
	projectDir := findProjectRoot()
	logDir := filepath.Join(projectDir, "logs")
	logFile := "dnslog_for_go.log_write"

	logFilePath := filepath.Join(logDir, logFile)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

}
