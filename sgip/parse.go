// 从缓冲区 读取解析sgip相关结构包

package sgip

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/yedamao/MockSMG/common"
	"io"
	"log"
)

func ParseHeader(buf io.Reader) (Head, error) {
	var header Head

	err := binary.Read(buf, binary.BigEndian, &header)
	if err != nil {
		return header, err
	}

	log.Println(header)
	return header, nil
}

func ParseBind(buf io.Reader) (Bind, error) {
	var bind Bind

	err := binary.Read(buf, binary.BigEndian, &bind)
	if err != nil {
		return bind, err
	}
	return bind, nil
}

func ParseSubmit(buf io.Reader, Total_len uint32) (Submit, error) {
	var submit Submit
	if err := binary.Read(buf, binary.BigEndian, &submit); err != nil {
		return submit, err
	}

	// parse msg
	msg := make([]byte, submit.MessageLength)
	if err := binary.Read(buf, binary.BigEndian, msg); err != nil {
		return submit, errors.New("parse msg error: " + err.Error())
	}
	msg, _ = common.Decodegbk(msg)
	fmt.Println(string(msg))

	reverse := make([]byte, 8)
	if err := binary.Read(buf, binary.BigEndian, reverse); err != nil {
		return submit, errors.New("parse reverse error: " + err.Error())
	}

	return submit, nil
}

func ParseResp(buf io.Reader) (Resp, error) {
	var resp Resp

	err := binary.Read(buf, binary.BigEndian, &resp)
	if err != nil {
		return resp, err
	}
	return resp, err
}
