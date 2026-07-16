package go_epay

import (
	"fmt"
	"os"
	"os/exec"
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
		FpxNotifyURL:      FpxNotifyURL,
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
		BankName:    "CIBBMYKL",
		Name:        "Goods",
		Money:       "100.00",
	}
	rsp, err := cli.Deposit(req)
	if err != nil {
		t.Logf("deposit err: %v", err)
		return
	}
	t.Logf("deposit rsp: code=%d tradeNo=%s payurl=%s", rsp.Code, rsp.TradeNo, rsp.PayURL)

	// 将 FormHTML 写入临时文件并用浏览器打开，验证自动 POST 跳转效果
	htmlFile := "/tmp/epay_deposit_test.html"
	if err := os.WriteFile(htmlFile, []byte(rsp.FormHTML), 0644); err != nil {
		t.Logf("write html file err: %v", err)
		return
	}
	t.Logf("FormHTML written to %s, opening in browser...", htmlFile)
	if err := exec.Command("open", htmlFile).Start(); err != nil {
		t.Logf("open browser err: %v", err)
	}
}
