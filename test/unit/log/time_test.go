package log

import (
	"fmt"
	"testing"
	"time"
)

func Test_log_time(t *testing.T) {

	fmt.Println(time.Now().Format("2006-01-02"))
}
