// mo packge

// 连接SP server
// 提交 状态报告 和 上行短信

package mo

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/yedamao/MockSMG/sgip"
)

var serial uint32 = 0

// 通过 pipe 与 mt packge 传递短信Head
func Run(pipe <-chan sgip.Head) {
	conn, err := net.Dial("tcp", "127.0.0.1:8002")
	if err != nil {
		log.Fatal("Dail SP error: ", err)
	}

	// 接收响应
	go recvResp(conn)

	if err := bindSP(conn); err != nil {
		log.Fatal(err)
	}

	// 接收head 发送回执状态
	for head := range pipe {
		fmt.Println(head)
	}
}

// 接收响应
func recvResp(conn net.Conn) {
	buf := bufio.NewReader(conn)
	log.Println("recv routine running")

	for {
		// 读取 包头
		resp, err := sgip.ParseResp(buf)
		if err != nil {
			if io.EOF == err {
				log.Println("EOF return", err)
				return
			}

			log.Println("parse resp error: ", err)
		}
		log.Println(resp)

		switch resp.Header.CMD {
		case sgip.SGIP_BIND_REP:
			log.Println("登陆响应")
		}
		log.Printf("seq %10d%10d%10d Result %d\n",
			resp.Header.Seq1, resp.Header.Seq2, resp.Header.Seq3, resp.Result)
	}
}

// 登陆SP
func bindSP(conn net.Conn) error {
	//  TODO 登陆SP

	bind := sgip.Bind{}
	copy(bind.Name[:], []byte("10690090"))
	copy(bind.Password[:], []byte("kjhhhg"))
	bind.Type = 2

	head := sgip.Head{}
	head.CMD = sgip.SGIP_BIND
	head.Seq1 = 3020092008
	head.Seq2 = getTimeStamp()
	head.Seq3 = getSerial()
	head.MsgLen = sgip.SGIP_BIND_LEN

	err := binary.Write(conn, binary.BigEndian, &head)
	if err != nil {
		return err
	}
	err = binary.Write(conn, binary.BigEndian, &bind)
	if err != nil {
		return err
	}
	return nil
}

func getTimeStamp() uint32 {
	t := time.Now()
	return uint32(int(t.Month())*100000000 + t.Day()*1000000 + t.Hour()*10000 + t.Minute()*100 + t.Second())
}

func getSerial() uint32 {
	serial += 1
	return serial
}
