// sgip 相关公用函数

package sgip

import (
	"time"
)

// 包序列号
var serial uint32 = 0

func getTimeStamp() uint32 {
	t := time.Now()
	return uint32(int(t.Month())*100000000 + t.Day()*1000000 +
		t.Hour()*10000 + t.Minute()*100 + t.Second())
}

func getSerial() uint32 {
	serial += 1
	return serial
}

func getICPId() uint32 {
	return 20002
	// return 3020092008
}

// 返回head, cmd, len 未赋值
func genHead() Head {
	head := Head{}
	head.Seq1 = getICPId()
	head.Seq2 = getTimeStamp()
	head.Seq3 = getSerial()

	return head
}
