package go_epay

import (
	"fmt"
	"testing"
	"time"

	"github.com/listenfengyang/go-epay/utils"
)

type testLogger struct{}

func (l *testLogger) Debugf(f string, a ...interface{}) { fmt.Printf("[DEBUG] "+f+"\n", a...) }
func (l *testLogger) Infof(f string, a ...interface{})  { fmt.Printf("[INFO]  "+f+"\n", a...) }
func (l *testLogger) Warnf(f string, a ...interface{})  { fmt.Printf("[WARN]  "+f+"\n", a...) }
func (l *testLogger) Errorf(f string, a ...interface{}) { fmt.Printf("[ERROR] "+f+"\n", a...) }

var _ utils.Logger = (*testLogger)(nil)

func newTestClient() *Client {
	return NewClient(&testLogger{}, &EPayInitParams{
		BaseURL:           DefaultBaseURL,
		MerchantID:        TestMerchantID,
		DepositKey:        DepositKey,
		PayoutKey:         PayoutKey,
		NotifyURL:         NotifyURL,
		ReturnURL:         ReturnURL,
		WithdrawNotifyURL: WithdrawNotifyURL,
	})
}

func TestDeposit(t *testing.T) {
	cli := newTestClient()
	cli.SetDebugModel(false)
	orderNo := fmt.Sprintf("T%d", time.Now().Unix())
	req := EPayDepositReq{
		Type:        "Bank",
		OutTradeNo:  orderNo,
		AccountName: "WangWei",
		BankAccount: "1234567890",
		BankName:    "Maybank",
		Name:        "Goods",
		Money:       "100.00",
		NotifyURL:   "http://notify.test.com/notify",
		ReturnURL:   "http://return.test.com/return",
	}
	rsp, err := cli.Deposit(req)
	if err != nil {
		t.Logf("deposit err: %v", err)
		return
	}
	t.Logf("deposit rsp: code=%d tradeNo=%s payurl=%s", rsp.Code, rsp.TradeNo, rsp.PayURL)
}
