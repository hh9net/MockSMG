// 向连接中 发送 sgip 相关包 函数

package sgip

import (
	"encoding/binary"
	"net"
)

// 返回应答包
func Response(conn net.Conn, head Head, code uint8) error {
	// 返回响应状态
	var resp Resp
	resp.Header = head
	resp.Header.MsgLen = SGIP_REP_LEN
	resp.Header.CMD += 0x80000000 // CMD 转换
	resp.Result = code

	return binary.Write(conn, binary.BigEndian, &resp)
}

// 向SP提交状态报告
func SubmitReport(conn net.Conn, head Head) error {
	report := Report{}
	report.Header = genHead()
	report.Header.CMD = SGIP_REPORT
	report.Header.MsgLen = SGIP_REPORT_LEN

	report.SubmitSequenceNumber[0] = head.Seq1
	report.SubmitSequenceNumber[1] = head.Seq2
	report.SubmitSequenceNumber[2] = head.Seq3

	report.ReportType = 0 // 0：对先前一条Submit命令的状态报告
	copy(report.UserNumber[:], []byte("18599999999"))
	report.State = 0 // 0：发送成功
	report.ErrorCode = 0

	return binary.Write(conn, binary.BigEndian, &report)
}

// 登陆SP
func BindSP(conn net.Conn) error {
	//  TODO 登陆SP

	bind := Bind{}
	copy(bind.Name[:], []byte("10690090"))
	copy(bind.Password[:], []byte("kjhhhg"))
	bind.Type = 2

	head := genHead()
	head.CMD = SGIP_BIND
	head.MsgLen = SGIP_BIND_LEN

	err := binary.Write(conn, binary.BigEndian, &head)
	if err != nil {
		return err
	}
	return binary.Write(conn, binary.BigEndian, &bind)
}
