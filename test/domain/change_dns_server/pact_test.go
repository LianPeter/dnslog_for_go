package change_dns_server

import (
	"fmt"
	"gopkg.in/ini.v1"
	"testing"
)

func TestIni(t *testing.T) {
	cfg, err := ini.Load("default.ini")
	if err != nil {
		panic("Unable to read configuration file")
	}

	current1 := cfg.Section("PACT").Key("udp").String()
	current2 := cfg.Section("PACT").Key("tcp").String()
	fmt.Println(current1)
	fmt.Println(current2)
}
