package go_epay

const (
	DefaultBaseURL = "https://api.epay365.biz"
	DepositPath    = "/submit.php" // Hosted Page 入金
	DepositViaPath = "/mapi.php"   // Server-to-Server API 入金
	WithdrawPath   = "/payout.php" // 出金

	// 测试用商户配置（生产环境通过 EPayInitParams 传入）
	// TestMerchantID = "Z100247"
	// DepositKey     = "uAiCJqmp52u5pf5kpEWWiJcSAB9B3WYx"
	// PayoutKey      = "i5MRR7r6mSmdv7OJBrYTqyL6BnYM7Q2O"

	TestMerchantID = "Z100257"
	DepositKey     = "bYjOdGxAk4icEhMnph8z2XdQhXSGUVph"
	PayoutKey      = "2Y9gmnFOUOhi3mVTE7KMLK4wx4OWzWas"

	NotifyURL         = "https://api-test.logtec.dev/fapi/v2/payment/psp/public/epay/deposit"
	WithdrawNotifyURL = "https://api-test.logtec.dev/fapi/v2/payment/psp/public/epay/withdraw"
	ReturnURL         = "https://portal.cptmarkets.com/"
)
