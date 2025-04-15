package unit

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func parseArgs(args []string) (domain, wordlist, server string, workers int) {
	// 设置默认值
	domain = "baidu.com"
	wordlist = ""
	server = "8.8.8.8:53"
	workers = 100

	// 简单的参数解析
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-domain":
			if i+1 < len(args) {
				domain = args[i+1]
				i++
			}
		case "-wordlist":
			if i+1 < len(args) {
				wordlist = args[i+1]
				i++
			}
		case "-server":
			if i+1 < len(args) {
				server = args[i+1]
				i++
			}
		case "-c":
			if i+1 < len(args) {
				if v, err := strconv.Atoi(args[i+1]); err == nil {
					workers = v
				}
				i++
			}
		}
	}
	return
}

func TestUseArgs(t *testing.T) {
	// 模拟命令行参数
	os.Args = []string{"test", "-domain", "example.com", "-wordlist", "dict.txt", "-server", "1.1.1.1:53", "-c", "50"}

	// 去掉第一个参数（程序名）
	args := os.Args[1:]

	// 解析参数
	domain, wordlist, server, workers := parseArgs(args)

	// 输出测试值
	fmt.Println("domain:", domain)
	fmt.Println("wordlist:", wordlist)
	fmt.Println("server:", server)
	fmt.Println("workers:", workers)

	// 测试断言
	if domain != "example.com" {
		t.Errorf("Expected domain to be example.com, got %s", domain)
	}
}
