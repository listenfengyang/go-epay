package go_epay

// ============================================================
// Pay-in (Deposit) - Hosted Page
// POST https://api.epay365.biz/submit.php
// ============================================================

// EPayDepositReq 入金请求参数
type EPayDepositReq struct {
	PID         string `mapstructure:"pid"          json:"pid"          form:"pid"`          // 商户 ID（自动填充）
	Type        string `mapstructure:"type"         json:"type"         form:"type"`         // 支付类型，可选，留空则跳转支付页
	OutTradeNo  string `mapstructure:"out_trade_no" json:"out_trade_no" form:"out_trade_no"` // 商户唯一订单号
	AccountName string `mapstructure:"account_name" json:"account_name" form:"account_name"` // 付款方账户持有人姓名
	BankAccount string `mapstructure:"bank_account" json:"bank_account" form:"bank_account"` // 付款方银行账户号
	BankName    string `mapstructure:"bank_name"    json:"bank_name"    form:"bank_name"`    // 收款银行名称
	NotifyURL   string `mapstructure:"notify_url"   json:"notify_url"   form:"notify_url"`   // 异步回调地址（自动填充）
	ReturnURL   string `mapstructure:"return_url"   json:"return_url"   form:"return_url"`   // 支付完成跳转地址（自动填充）
	Name        string `mapstructure:"name"         json:"name"         form:"name"`         // 商品名称
	Money       string `mapstructure:"money"        json:"money"        form:"money"`        // 金额，单位 MYR，保留两位小数
	Param       string `mapstructure:"param"        json:"param"        form:"param"`        // 透传参数，可选
}

// EPayDepositRsp 入金响应
type EPayDepositRsp struct {
	Code      int    `json:"code"      form:"code"`      // 1=成功，其他失败
	Msg       string `json:"msg"       form:"msg"`       // 消息
	TradeNo   string `json:"trade_no"  form:"trade_no"`  // 平台订单号
	PayURL    string `json:"payurl"    form:"payurl"`    // 支付提交地址（POST action URL）
	QRCode    string `json:"qrcode"    form:"qrcode"`    // 二维码链接
	URLScheme string `json:"urlscheme" form:"urlscheme"` // 小程序 scheme
	Param     string `json:"param"     form:"param"`     // 透传参数，如银行卡信息
	FormHTML  string `json:"form_html" form:"form_html"` // 自动提交的 HTML 表单，嵌入页面后浏览器自动 POST 到支付页
}

// ============================================================
// Pay-out (Withdraw)
// ============================================================

// EPayWithdrawReq 出金请求参数
// POST https://api.epay365.biz/payout.php
type EPayWithdrawReq struct {
	PID         string `mapstructure:"pid"          json:"pid"          form:"pid"`          // 商户 ID（自动填充）
	Type        string `mapstructure:"type"         json:"type"         form:"type"`         // 出金类型: 1=personal, 2=corporate
	OutTradeNo  string `mapstructure:"out_trade_no" json:"out_trade_no" form:"out_trade_no"` // 商户唯一订单号
	NotifyURL   string `mapstructure:"notify_url"   json:"notify_url"   form:"notify_url"`   // 出金回调地址（自动填充）
	Name        string `mapstructure:"name"         json:"name"         form:"name"`         // 收款人全名
	BankID      string `mapstructure:"bank_id"      json:"bank_id"      form:"bank_id"`      // 银行编码（见 Bank Codes Table）
	BankAccount string `mapstructure:"bank_account" json:"bank_account" form:"bank_account"` // 收款银行账户
	Branch      string `mapstructure:"branch"       json:"branch"       form:"branch"`       // 支行名称
	Money       string `mapstructure:"money"        json:"money"        form:"money"`        // 金额，单位 MYR，格式: 100.00
	ClientIP    string `mapstructure:"clientip"     json:"clientip"     form:"clientip"`     // 发起方 IP
	Param       string `mapstructure:"param"        json:"param"        form:"param"`        // 业务扩展参数（可选）
}

// EPayWithdrawRsp 出金响应
type EPayWithdrawRsp struct {
	Code    int    `json:"code"     form:"code"`     // 1=成功，其他失败
	Msg     string `json:"msg"      form:"msg"`      // 消息
	TradeNo string `json:"trade_no" form:"trade_no"` // 平台订单号
}

// ============================================================
// Callback（入金 & 出金回调）
// ============================================================

// EPayCallbackReq 回调通知参数
type EPayCallbackReq struct {
	PID         string `json:"pid"          form:"pid"`          // 商户 ID
	TradeNo     string `json:"trade_no"     form:"trade_no"`     // 平台订单号
	OutTradeNo  string `json:"out_trade_no" form:"out_trade_no"` // 商户订单号
	Type        string `json:"type"         form:"type"`         // 支付类型
	Name        string `json:"name"         form:"name"`         // 商品名称
	Money       string `json:"money"        form:"money"`        // 金额
	TradeStatus string `json:"trade_status" form:"trade_status"` // 交易状态：success=成功
	Param       string `json:"param"        form:"param"`        // 透传参数
	Sign        string `json:"sign"         form:"sign"`         // 签名
	SignType    string `json:"sign_type"    form:"sign_type"`    // 签名类型，固定 MD5
}

// ============================================================
// Query
// ============================================================

// EPayQueryOrderRsp 查询订单响应
type EPayQueryOrderRsp struct {
	Code        int     `json:"code"         form:"code"`         // 1=成功，其他失败
	Msg         string  `json:"msg"          form:"msg"`          // 消息
	PID         string  `json:"pid"          form:"pid"`          // 商户 ID
	TradeNo     string  `json:"trade_no"     form:"trade_no"`     // 平台订单号
	OutTradeNo  string  `json:"out_trade_no" form:"out_trade_no"` // 商户订单号
	Type        string  `json:"type"         form:"type"`         // 支付类型
	Money       float64 `json:"money"        form:"money"`        // 金额
	TradeStatus string  `json:"trade_status" form:"trade_status"` // 交易状态：success=成功
}

// EPayDepositViaReq Create Payment via API 请求参数
type EPayDepositViaReq struct {
	PID               string `mapstructure:"pid"                 json:"pid"                 form:"pid"`                 // 商户 ID（自动填充）
	Type              string `mapstructure:"type"                json:"type"                form:"type"`                // 支付类型
	OutTradeNo        string `mapstructure:"out_trade_no"        json:"out_trade_no"        form:"out_trade_no"`        // 商户唯一订单号
	AccountName       string `mapstructure:"account_name"        json:"account_name"        form:"account_name"`        // 付款方账户持有人姓名
	BankAccount       string `mapstructure:"bank_account"        json:"bank_account"        form:"bank_account"`        // 付款方银行账户号
	BankName          string `mapstructure:"bank_name"           json:"bank_name"           form:"bank_name"`           // 收款银行名称
	NotifyURL         string `mapstructure:"notify_url"          json:"notify_url"          form:"notify_url"`          // 异步回调地址（自动填充）
	ReturnURL         string `mapstructure:"return_url"          json:"return_url"          form:"return_url"`          // 支付完成跳转地址（自动填充）
	WithdrawNotifyUrl string `mapstructure:"withdraw_notify_url" json:"withdraw_notify_url" form:"withdraw_notify_url"` // 出金回调地址
	Name              string `mapstructure:"name"                json:"name"                form:"name"`                // 商品名称
	Money             string `mapstructure:"money"               json:"money"               form:"money"`               // 金额，单位 MYR，保留两位小数
	ClientIP          string `mapstructure:"clientip"            json:"clientip"            form:"clientip"`            // 发起方 IP
	Device            string `mapstructure:"device"              json:"device"              form:"device"`              // 设备类型，如 pc / mobile
	Param             string `mapstructure:"param"               json:"param"               form:"param"`               // 透传参数，可选
}

// EPayDepositViaRsp Create Payment via API 响应
type EPayDepositViaRsp struct {
	Code    int    `json:"code"     form:"code"`     // 1=成功，其他失败
	Msg     string `json:"msg"      form:"msg"`      // 消息
	TradeNo string `json:"trade_no" form:"trade_no"` // 平台订单号
	PayURL  string `json:"payurl"   form:"payurl"`   // 支付链接
	Param   string `json:"param"    form:"param"`    // 透传参数
}
