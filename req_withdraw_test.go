package go_epay

import (
	"fmt"
	"testing"
	"time"
)

func TestWithdraw(t *testing.T) {
	cli := newTestClient()
	cli.SetDebugModel(false)
	orderNo := fmt.Sprintf("WD%d", time.Now().Unix())
	req := EPayWithdrawReq{
		Type:        "1", // 1: personal, 2: corporate
		OutTradeNo:  orderNo,
		Name:        "Test User",
		BankID:      BankCodeMaybank, // MYRTIGER - Malayan Banking Berhad
		BankAccount: "1234567890",
		Branch:      "KL Main Branch",
		Money:       "50.00",
		ClientIP:    "18.162.184.178",
		NotifyURL:   WithdrawNotifyURL,
	}
	rsp, err := cli.Withdraw(req)
	if err != nil {
		t.Logf("withdraw err: %v", err)
		return
	}
	t.Logf("withdraw rsp: code=%d tradeNo=%s msg=%s", rsp.Code, rsp.TradeNo, rsp.Msg)
}
