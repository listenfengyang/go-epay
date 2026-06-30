package go_epay

// ============================================================
// Pay-in (Deposit) - Hosted Page
// POST https://api.epay365.biz/submit.php
// ============================================================

// EPayDepositReq 入金请求参数
type EPayDepositReq struct {
	PID         string `mapstructure:"pid"`          // 商户 ID（自动填充）
	Type        string `mapstructure:"type"`         // 支付类型，可选，留空则跳转支付页
	OutTradeNo  string `mapstructure:"out_trade_no"` // 商户唯一订单号
	AccountName string `mapstructure:"account_name"` // 付款方账户持有人姓名
	BankAccount string `mapstructure:"bank_account"` // 付款方银行账户号
	BankName    string `mapstructure:"bank_name"`    // 收款银行名称
	NotifyURL   string `mapstructure:"notify_url"`   // 异步回调地址（自动填充）
	ReturnURL   string `mapstructure:"return_url"`   // 支付完成跳转地址（自动填充）

	Name  string `mapstructure:"name"`  // 商品名称
	Money string `mapstructure:"money"` // 金额，单位 MYR，保留两位小数
	Param string `mapstructure:"param"` // 透传参数，可选
}

// EPayDepositRsp 入金响应
type EPayDepositRsp struct {
	Code      int    `json:"code"`      // 1=成功，其他失败
	Msg       string `json:"msg"`       // 消息
	TradeNo   string `json:"trade_no"`  // 平台订单号
	PayURL    string `json:"payurl"`    // 支付链接（与 qrcode/urlscheme 三选一）
	QRCode    string `json:"qrcode"`    // 二维码链接
	URLScheme string `json:"urlscheme"` // 小程序 scheme
	Param     string `json:"param"`     // 透传参数，如银行卡信息
}

// ============================================================
// Pay-out (Withdraw)
// ============================================================

// EPayWithdrawReq 出金请求参数
// POST https://api.epay365.biz/payout.php
type EPayWithdrawReq struct {
	PID         string `mapstructure:"pid"`          // 商户 ID（自动填充）
	Type        string `mapstructure:"type"`         // 出金类型: 1=personal, 2=corporate
	OutTradeNo  string `mapstructure:"out_trade_no"` // 商户唯一订单号
	NotifyURL   string `mapstructure:"notify_url"`   // 出金回调地址（自动填充）
	Name        string `mapstructure:"name"`         // 收款人全名
	BankID      string `mapstructure:"bank_id"`      // 银行编码（见 Bank Codes Table）
	BankAccount string `mapstructure:"bank_account"` // 收款银行账户
	Branch      string `mapstructure:"branch"`       // 支行名称
	Money       string `mapstructure:"money"`        // 金额，单位 MYR，格式: 100.00
	ClientIP    string `mapstructure:"clientip"`     // 发起方 IP
	Param       string `mapstructure:"param"`        // 业务扩展参数（可选）
}

// EPayWithdrawRsp 出金响应
type EPayWithdrawRsp struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	TradeNo string `json:"trade_no"`
}

// ============================================================
// Callback（入金 & 出金回调）
// ============================================================

// EPayCallbackReq 回调通知参数
type EPayCallbackReq struct {
	PID         string `json:"pid"          form:"pid"`
	TradeNo     string `json:"trade_no"     form:"trade_no"`
	OutTradeNo  string `json:"out_trade_no" form:"out_trade_no"`
	Type        string `json:"type"         form:"type"`
	Name        string `json:"name"         form:"name"`
	Money       string `json:"money"        form:"money"`
	TradeStatus string `json:"trade_status" form:"trade_status"`
	Param       string `json:"param"        form:"param"`
	Sign        string `json:"sign"         form:"sign"`
	SignType    string `json:"sign_type"    form:"sign_type"`
}

// ============================================================
// Query
// ============================================================

// EPayQueryOrderRsp 查询订单响应
type EPayQueryOrderRsp struct {
	Code        int     `json:"code"`
	Msg         string  `json:"msg"`
	PID         string  `json:"pid"`
	TradeNo     string  `json:"trade_no"`
	OutTradeNo  string  `json:"out_trade_no"`
	Type        string  `json:"type"`
	Money       float64 `json:"money"`
	TradeStatus string  `json:"trade_status"`
}

// EPayDepositViaReq Create Payment via API 请求参数
type EPayDepositViaReq struct {
	PID               string `mapstructure:"pid"`
	Type              string `mapstructure:"type"`
	OutTradeNo        string `mapstructure:"out_trade_no"`
	AccountName       string `mapstructure:"account_name"`
	BankAccount       string `mapstructure:"bank_account"`
	BankName          string `mapstructure:"bank_name"`
	NotifyURL         string `mapstructure:"notify_url"`
	ReturnURL         string `mapstructure:"return_url"`
	WithdrawNotifyUrl string `mapstructure:"withdraw_notify_url"`
	Name              string `mapstructure:"name"`
	Money             string `mapstructure:"money"`
	ClientIP          string `mapstructure:"clientip"`
	Device            string `mapstructure:"device"`
	Param             string `mapstructure:"param"`
}

// EPayDepositViaRsp Create Payment via API 响应
type EPayDepositViaRsp struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	TradeNo string `json:"trade_no"`
	PayURL  string `json:"payurl"`
	Param   string `json:"param"`
}
