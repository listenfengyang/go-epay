package go_epay

// EPayInitParams epay 初始化参数
type EPayInitParams struct {
	BaseURL           string `json:"baseUrl"           mapstructure:"baseUrl"           config:"baseUrl"           yaml:"baseUrl"`           // API 根地址，默认 https://api.epay365.biz
	MerchantID        string `json:"merchantId"        mapstructure:"merchantId"        config:"merchantId"        yaml:"merchantId"`        // 商户 ID (pid)
	DepositKey        string `json:"depositKey"        mapstructure:"depositKey"        config:"depositKey"        yaml:"depositKey"`        // 入金密钥
	PayoutKey         string `json:"payoutKey"         mapstructure:"payoutKey"         config:"payoutKey"         yaml:"payoutKey"`         // 出金密钥
	NotifyURL         string `json:"notifyUrl"         mapstructure:"notifyUrl"         config:"notifyUrl"         yaml:"notifyUrl"`         // 入金异步回调地址
	ReturnURL         string `json:"returnUrl"         mapstructure:"returnUrl"         config:"returnUrl"         yaml:"returnUrl"`         // 入金同步跳转地址
	WithdrawNotifyURL string `json:"withdrawNotifyUrl" mapstructure:"withdrawNotifyUrl" config:"withdrawNotifyUrl" yaml:"withdrawNotifyUrl"` // 出金回调地址
}
