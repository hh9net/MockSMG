package main

// Command ID
const (
	SGIP_BIND = 0x00000001 + iota
	SGIP_UNBIND
	SGIP_SUBMIT
	SGIP_DELIVER
	SGIP_REPORT
	SGIP_ADDSP
	SGIP_MODIFYSP
	SGIP_DELETESP
	SGIP_QUERYROUTE
	SGIP_ADDTELESEG
	SGIP_MODIFYTELESEG
	SGIP_DELETETELESEG
	SGIP_ADDSMG
	SGIP_MODIFYSMG
	SGIP_DELETESMG
	SGIP_CHECKUSER
	SGIP_USERRPT
)

const (
	SGIP_BIND_REP = 0x80000001 + iota
	SGIP_UNBIND_REP
	SGIP_SUBMIT_REP
	SGIP_DELIVER_REP
	SGIP_REPORT_REP
	SGIP_ADDSP_REP
	SGIP_MODIFYSP_REP
	SGIP_DELETESP_REP
	SGIP_QUERYROUTE_REP
	SGIP_ADDTELESEG_REP
	SGIP_MODIFYTELESEG_REP
	SGIP_DELETETELESEG_REP
	SGIP_ADDSMG_REP
	SGIP_MODIFYSMG_REP
	SGIP_DELETESMG_REP
	SGIP_CHECKUSER_REP
	SGIP_USERRPT_REP
)

// 返回状态码
const (
	SUCC = iota
)

// MessageCoding
const (
	ASCII = 0  // 纯ASCII字符串
	UCS2  = 8  // UCS2编码
	GBK   = 15 // GBK编码
)

// package len
const (
	SGIP_HEAD_LEN = 20
	SGIP_REP_LEN  = SGIP_HEAD_LEN + 9
)

type Head struct {
	// Message Length 消息的总长度(字节)
	MsgLen uint32
	// Command ID 命令ID
	CMD uint32
	// Sequence Number 序列号
	Seq1 uint32
	Seq2 uint32
	Seq3 uint32
}

type Bind struct {
	Type     uint8    // 登录类型
	Name     [16]byte // 登陆名
	Password [16]byte // 密码
	Reverse  [8]byte  // 保留，扩展用
}

type Submit struct {
	SPNumber         [21]byte // SP的接入号码
	ChargeNumber     [21]byte // 付费号码，手机号码前加“86”国别标志
	UserCount        uint8    // 接收短消息的手机数量，取值范围1至100
	UserNumber       [21]byte // 接收该短消息的手机号
	CorpId           [5]byte  // 企业代码，取值范围0-99999
	ServiceType      [10]byte // 业务代码，由SP定义
	FeeType          uint8    // 业务代码，由SP定义
	FeeValue         [6]byte  // 该条短消息的收费值
	GivenValue       [6]byte  // 赠送用户的话费
	AgentFlag        uint8    // 代收费标志，0：应收；1：实收
	MorelatetoMTFlag uint8    // 引起MT消息的原因
	Priority         uint8    // 优先级0-9从低到高，默认为0
	ExpireTime       [16]byte
	ScheduleTime     [16]byte
	ReportFlag       uint8  // 状态报告标记
	TP_pid           uint8  // GSM协议类型
	TP_udhi          uint8  // GSM协议类型
	MessageCoding    uint8  // 短消息的编码格式
	MessageType      uint8  // 信息类型：
	MessageLength    uint32 // 短消息的长度
}

type Resp struct {
	Header  Head
	Result  uint8   // Bind执行命令是否成功
	Reverse [8]byte // 保留，扩展用
}
