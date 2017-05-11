package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net"
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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := bufio.NewReader(conn)

	for {
		// 读取 包头
		head, err := parseHeader(buf)
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

			bind, err := parseBind(buf)
			if err != nil {
				log.Println("parse bind pkg error: ", err)
				return
			}

			// to do  校验用户
			loginCheck(bind)

			err = response(conn, head, SUCC)
			if err != nil {
				log.Println("bind write resp error: ", err)
			}

			log.Println("login success")

		// 登陆包
		case SGIP_SUBMIT:
			log.Println("下发包")

			submit, err := parseSubmit(buf, head.MsgLen)
			if err != nil {
				log.Println("parse submit error: ", err)
				return
			}

			log.Println(string(submit.SPNumber[:16]))

			err = response(conn, head, SUCC)
			if err != nil {
				log.Println("submit write resp error: ", err)
			}
		default:
			log.Println("CMD not found: ", head.CMD)
		}

	}
}

func parseBind(buf io.Reader) (Bind, error) {
	var bind Bind

	err := binary.Read(buf, binary.BigEndian, &bind)
	if err != nil {
		return bind, err
	}
	return bind, nil
}

func parseHeader(buf io.Reader) (Head, error) {
	var header Head

	err := binary.Read(buf, binary.BigEndian, &header)
	if err != nil {
		return header, err
	}

	return header, nil
}

func parseSubmit(buf io.Reader, Total_len uint32) (Submit, error) {
	var submit Submit
	if err := binary.Read(buf, binary.BigEndian, &submit); err != nil {
		return submit, err
	}

	// parse msg
	msg := make([]byte, submit.MessageLength)
	if err := binary.Read(buf, binary.BigEndian, msg); err != nil {
		log.Println("parse msg error: ", err)
	}

	reverse := make([]byte, 8)
	if err := binary.Read(buf, binary.BigEndian, reverse); err != nil {
		log.Println("parse reverse error: ", err)
	}

	if UCS2 == submit.MessageCoding {
		log.Println(string(msg))
	}
	if GBK == submit.MessageCoding {
		msg, err := Decodegbk(msg)
		if err != nil {
			log.Println("convert gbk to utf-8 error: ", err)
		}
		log.Println(string(msg))
	}

	return submit, nil
}

// 校验用户
func loginCheck(bind Bind) bool {
	// ToDo
	return true
}

// 返回应答包
func response(conn net.Conn, head Head, code uint8) error {
	// 返回响应状态
	var resp Resp
	resp.Header = head
	resp.Header.MsgLen = SGIP_REP_LEN
	resp.Header.CMD += 0x80000000
	resp.Result = code

	return binary.Write(conn, binary.BigEndian, &resp)
}

//convert GBK to UTF-8
func Decodegbk(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e

	}
	return d, nil
}
