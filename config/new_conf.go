package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type DnsConfig struct {
	Domain string `yaml:"domain"`
}

// NewYaml 生成yaml配置文件
func NewYaml(domainName string) error {

	config := DnsConfig{}
	config.Domain = domainName

	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Println("Error marshalling config to YAML:", err)
		return err
	}

	err = os.WriteFile("config.yaml", data, 0644)
	if err != nil {
		log.Println("Error writing config.yaml:", err)
		return err
	}

	if domainName == "" {
		log.Println("域名为空，无法创建配置文件")
		return fmt.Errorf("domain name is empty")
	}

	return nil
}
