package go_epay

// BankCode 银行/支付渠道编码结构体
// 支持多币种的银行列表映射
type BankCode struct {
	Currency       string `json:"currency"`       // 币种，如 MYR、PHP
	Channel        string `json:"channel"`        // 支付通道
	Code           string `json:"code"`           // 银行/渠道编码（epay 侧）
	Name           string `json:"name"`           // 银行/渠道名称
	ClientBankCode string `json:"clientBankCode"` // 客户侧银行编码
}

// BankCodes 全币种银行列表
var BankCodes = []BankCode{
	// -------------------------------------------------------
	// MYR - Malaysia Ringgit
	// -------------------------------------------------------
	{Currency: "MYR", Channel: "epay", Code: "MYRTIGER", Name: "Malayan Banking Berhad", ClientBankCode: "MBB"},
	{Currency: "MYR", Channel: "epay", Code: "CIBBMYKL", Name: "CIMB Bank", ClientBankCode: "CIMB"},
	{Currency: "MYR", Channel: "epay", Code: "PBBEMYKL", Name: "Public Bank", ClientBankCode: "PBB"},
	{Currency: "MYR", Channel: "epay", Code: "HLBBMYKL", Name: "Hong Leong Bank", ClientBankCode: "HLB"},
	{Currency: "MYR", Channel: "epay", Code: "HBMBMYKL", Name: "HSBC Bank", ClientBankCode: "HSBC"},
	{Currency: "MYR", Channel: "epay", Code: "OCBCMYKL", Name: "OCBC Bank", ClientBankCode: "OCBC"},
	{Currency: "MYR", Channel: "epay", Code: "RHBBMYKL", Name: "RHB Bank", ClientBankCode: "RHB"},
	{Currency: "MYR", Channel: "epay", Code: "ARBKMYKL", Name: "AmBank", ClientBankCode: "AMB"},
	{Currency: "MYR", Channel: "epay", Code: "PHBMMYKL", Name: "Affin Bank", ClientBankCode: "AFF"},
	{Currency: "MYR", Channel: "epay", Code: "MFBBMYKL", Name: "Alliance Bank", ClientBankCode: "ALB"},
	{Currency: "MYR", Channel: "epay", Code: "SCBLMYKX", Name: "Standard Chartered Bank", ClientBankCode: "SCTB"},
	{Currency: "MYR", Channel: "epay", Code: "UOVBMYKL", Name: "United Overseas Bank", ClientBankCode: "UOB"},
	{Currency: "MYR", Channel: "epay", Code: "BIMBMYKL", Name: "Bank Islam", ClientBankCode: ""},
	{Currency: "MYR", Channel: "epay", Code: "BMMBMYKL", Name: "Bank Muamalat", ClientBankCode: "BMMBMYKL"},
	{Currency: "MYR", Channel: "epay", Code: "BKRMMYKL", Name: "Bank Rakyat", ClientBankCode: ""},
	{Currency: "MYR", Channel: "epay", Code: "BSNAMYK1", Name: "Bank Simpanan Nasional", ClientBankCode: "BSN"},
	{Currency: "MYR", Channel: "epay", Code: "GXSPMYKL", Name: "GX Bank", ClientBankCode: ""},
	{Currency: "MYR", Channel: "epay", Code: "TNGDMYNB", Name: "Touch N Go", ClientBankCode: ""},
	{Currency: "MYR", Channel: "epay", Code: "BOSTMYNB", Name: "Boost", ClientBankCode: ""},
}

// -------------------------------------------------------
// 向后兼容：保留 MYR 银行编码常量，供出金请求使用
// -------------------------------------------------------

// MYR 马来西亚银行编码常量（用于出金请求 bank_id 字段）
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
