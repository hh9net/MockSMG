// mt package
// 接收sp请求, 返回resp
// 将submit 请求的 head 通过pipe 传递给mo返回提交状态

package mt

import (
	"bufio"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	. "github.com/yedamao/MockSMG/sgip"
)

func Run(pipe chan<- Head) {
	ln, err := net.Listen("tcp", ":8801")
	if err != nil {
		log.Fatal("Listen 8801 error: ", err)
	}
	defer ln.Close()
	log.Println("Listen 8801  .....")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error: ", err)
		}
		go handleSPCli(conn, pipe)
	}
}

// 接收响应
func recvResp(conn net.Conn) {
	buf := bufio.NewReader(conn)
	log.Println("recv routine running")

	for {
		// 读取 包头
		resp, err := ParseResp(buf)
		if err != nil {
			if io.EOF == err {
				log.Println("EOF return", err)
				return
			}

			log.Println("parse resp error: ", err)
		}
		log.Println(resp)

		switch resp.Header.CMD {
		case SGIP_BIND_REP:
			log.Println("登陆响应")
		}
		log.Printf("seq %10d%10d%10d Result %d\n",
			resp.Header.Seq1, resp.Header.Seq2, resp.Header.Seq3, resp.Result)
	}
}

// 处理sp请求
func handleSPCli(conn net.Conn, pipe chan<- Head) {
	defer conn.Close()

	buf := bufio.NewReader(conn)

	for {
		// 读取 包头
		head, err := ParseHeader(buf)
		if err != nil {
			if io.EOF == err {
				log.Println("EOF return", err)
				return
			}

			log.Println("parse header error: ", err)
		}

		switch head.CMD {
		case SGIP_BIND:
			log.Println("登陆包")

			bind, err := ParseBind(buf)
			if err != nil {
				log.Println("parse bind pkg error: ", err)
				return
			}

			// to do  校验用户
			loginCheck(bind)

			err = Response(conn, head, SUCC)
			if err != nil {
				log.Println("bind write resp error: ", err)
			}

			log.Println("login success")

		// 登陆包
		case SGIP_SUBMIT:
			log.Println("下发包")

			submit, err := ParseSubmit(buf, head.MsgLen)
			if err != nil {
				log.Println("parse submit error: ", err)
				return
			}

			log.Println(string(submit.SPNumber[:16]))

			// 返回响应
			go goResp(conn, head, SUCC)
			// 传递 head给mo 返回状态报告
			pipe <- head

		default:
			log.Println("CMD not found: ", head.CMD)
		}
	}
}

// 校验用户
func loginCheck(bind Bind) bool {
	// ToDo
	return true
}

func goResp(conn net.Conn, head Head, result int) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	err := Response(conn, head, SUCC)
	if err != nil {
		log.Println("submit write resp error: ", err)
	}
}
