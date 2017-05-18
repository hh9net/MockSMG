// 联通网关相关

// 相关包 结构体定义

package sgip

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
	Type     uint8    // 登录类型 1 sp -> SMG, 2 SMG -> SP
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

// SMG用Deliver命令向SP发送一条MO短消息
type Deliver struct {
	Header        Head
	UserNumber    [21]byte // 接收该短消息的手机号
	SPNumber      [21]byte // SP的接入号码
	TP_pid        uint8    // GSM协议类型
	TP_udhi       uint8    // GSM协议类型
	MessageCoding uint8    // 短消息的编码格式
	MessageLength uint32   // 短消息的长度
	// MessageContent
	// Reverse
}

// Report命令用于向SP发送一条先前的Submit命令的当前状态
type Report struct {
	Header               Head
	SubmitSequenceNumber [3]uint32
	ReportType           uint8
	UserNumber           [21]byte // 接收该短消息的手机号
	State                uint8    // 0：发送成功 1：等待发送 2：发送失败
	ErrorCode            uint8    // 当State=2时为错误码值，否则为0
	Reserve              [8]byte
}
