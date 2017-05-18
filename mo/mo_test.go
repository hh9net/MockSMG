package mo

import (
	"github.com/yedamao/MockSMG/sgip"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	pipe := make(chan sgip.Head)
	go mockSend(pipe)
	Run(pipe)
}

// 每隔一秒向管道发送head
func mockSend(pipe chan<- sgip.Head) {
	for i := 0; i < 10; i++ {
		pipe <- sgip.Head{
			sgip.SGIP_HEAD_LEN,
			sgip.SGIP_BIND_REP,
			20002,
			5191110,
			getSerial(),
		}

		time.Sleep(time.Second)
	}
	close(pipe)
}
