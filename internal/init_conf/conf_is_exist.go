package init_conf

import (
	"dnslog_for_go/internal/config"
	"log"
	"os"
)

// IsExist 如果配置存在则不进行创建
func IsExist() error {
	_, err := os.Stat("internal/config/domain.yaml")
	if err == nil { // 文件存在
		log.Println("config.yaml already exists, no need to create a new one.")
		return nil
	} else if os.IsNotExist(err) { // 文件不存在
		log.Println("config.yaml not found, creating new configuration...")
		// 创建配置文件
		err := NewYaml(config.GlobalDomainNameForGetDomain)
		if err != nil { // 创建失败
			log.Println("Error creating config.yaml:", err)
			return err
		}
		log.Println("config.yaml created successfully")
		return nil
	} else { // 其他错误
		return err
	}
}
