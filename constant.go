package go_epay

const (
	DefaultBaseURL = "https://api.epay365.biz"
	DepositPath    = "/submit.php" // Hosted Page 入金
	DepositViaPath = "/mapi.php"   // Server-to-Server API 入金
	WithdrawPath   = "/payout.php" // 出金

	// 以下为测试用商户配置，本地填写真实值，不提交到 git
	TestMerchantID = ""
	DepositKey     = ""
	PayoutKey      = ""

	NotifyURL         = ""
	WithdrawNotifyURL = ""
	ReturnURL         = ""
)
