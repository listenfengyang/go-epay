package go_epay

import (
	"fmt"
	"testing"
	"time"
)

func TestDepositVia(t *testing.T) {
	cli := newTestClient()
	cli.SetDebugModel(false)
	orderNo := fmt.Sprintf("TV%d", time.Now().Unix())
	req := EPayDepositViaReq{
		Type:        "Bank",
		OutTradeNo:  orderNo,
		AccountName: "WangWei",
		BankAccount: "1234567890",
		BankName:    "Maybank",
		Name:        "Goods",
		Money:       "100.00",
		ClientIP:    "18.162.184.178",
		NotifyURL:   NotifyURL,
		ReturnURL:   ReturnURL,
	}
	rsp, err := cli.DepositVia(req)
	if err != nil {
		t.Logf("deposit via err: %v", err)
		return
	}
	t.Logf("deposit via rsp: code=%d tradeNo=%s payurl=%s param=%s", rsp.Code, rsp.TradeNo, rsp.PayURL, rsp.Param)
}
