package main

import (
	"bufio"
	"io"
	"log"
	"net"

	. "github.com/yedamao/MockSMG/sgip"
)

func main() {
	ln, err := net.Listen("tcp", ":8801")
	if err != nil {
		log.Fatal("Listen 8801 error: ", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error: ", err)
		}
		go handleSPCli(conn)
	}
}

// 处理sp请求
func handleSPCli(conn net.Conn) {
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

			err = Response(conn, head, SUCC)
			if err != nil {
				log.Println("submit write resp error: ", err)
			}
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
