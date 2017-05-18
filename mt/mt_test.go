package mt

import (
	"github.com/yedamao/MockSMG/sgip"
	"log"
	"testing"
)

func TestRun(t *testing.T) {
	pipe := make(chan sgip.Head)
	go readPipe(pipe)
	Run(pipe)
}

func readPipe(pipe <-chan sgip.Head) {
	for head := range pipe {
		log.Println(head)
	}
}
