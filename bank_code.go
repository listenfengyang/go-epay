package go_epay

// BankCode 马来西亚银行编码常量
// 用于出金请求中的 bank_id 字段
const (
	BankCodeMaybank         = "MYRTIGER" // Malayan Banking Berhad (Maybank)
	BankCodeCIMB            = "CIBBMYKL" // CIMB Bank
	BankCodePublicBank      = "PBBEMYKL" // Public Bank
	BankCodeHongLeong       = "HLBBMYKL" // Hong Leong Bank
	BankCodeHSBC            = "HBMBMYKL" // HSBC Bank
	BankCodeOCBC            = "OCBCMYKL" // OCBC Bank
	BankCodeRHB             = "RHBBMYKL" // RHB Bank
	BankCodeAmBank          = "ARBKMYKL" // AmBank
	BankCodeAffin           = "PHBMMYKL" // Affin Bank
	BankCodeAlliance        = "MFBBMYKL" // Alliance Bank
	BankCodeStandardCharter = "SCBLMYKX" // Standard Chartered Bank
	BankCodeUOB             = "UOVBMYKL" // United Overseas Bank
	BankCodeBankIslam       = "BIMBMYKL" // Bank Islam
	BankCodeBankMuamalat    = "BMMBMYKL" // Bank Muamalat
	BankCodeBankRakyat      = "BKRMMYKL" // Bank Rakyat
	BankCodeBSN             = "BSNAMYK1" // Bank Simpanan Nasional
	BankCodeGXBank          = "GXSPMYKL" // GX Bank
	BankCodeTouchNGo        = "TNGDMYNB" // Touch N Go
	BankCodeBoost           = "BOSTMYNB" // Boost
)

// BankCodeNames 银行编码与名称映射表
var BankCodeNames = map[string]string{
	BankCodeMaybank:         "Malayan Banking Berhad",
	BankCodeCIMB:            "CIMB Bank",
	BankCodePublicBank:      "Public Bank",
	BankCodeHongLeong:       "Hong Leong Bank",
	BankCodeHSBC:            "HSBC Bank",
	BankCodeOCBC:            "OCBC Bank",
	BankCodeRHB:             "RHB Bank",
	BankCodeAmBank:          "AmBank",
	BankCodeAffin:           "Affin Bank",
	BankCodeAlliance:        "Alliance Bank",
	BankCodeStandardCharter: "Standard Chartered Bank",
	BankCodeUOB:             "United Overseas Bank",
	BankCodeBankIslam:       "Bank Islam",
	BankCodeBankMuamalat:    "Bank Muamalat",
	BankCodeBankRakyat:      "Bank Rakyat",
	BankCodeBSN:             "Bank Simpanan Nasional",
	BankCodeGXBank:          "GX Bank",
	BankCodeTouchNGo:        "Touch N Go",
	BankCodeBoost:           "Boost",
}
