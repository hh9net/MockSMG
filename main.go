package main

import (
	"github.com/yedamao/MockSMG/mo"
	"github.com/yedamao/MockSMG/mt"
	"github.com/yedamao/MockSMG/sgip"
)

func main() {
	pipe := make(chan sgip.Head)

	go mt.Run(pipe)

	mo.Run(pipe)
}
