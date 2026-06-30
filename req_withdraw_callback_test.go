package go_epay

import (
	"testing"

	"github.com/listenfengyang/go-epay/utils"
)

func TestWithdrawCallback(t *testing.T) {
	cli := newTestClient()
	params := map[string]string{
		"pid":          TestMerchantID,
		"trade_no":     "EPAY20240101002",
		"out_trade_no": "WD_ORDER_001",
		"type":         "withdraw",
		"name":         "Withdraw",
		"money":        "50.00",
		"trade_status": "TRADE_SUCCESS",
	}
	sign := utils.Sign(params, PayoutKey)
	req := EPayCallbackReq{
		PID:         params["pid"],
		TradeNo:     params["trade_no"],
		OutTradeNo:  params["out_trade_no"],
		Type:        params["type"],
		Name:        params["name"],
		Money:       params["money"],
		TradeStatus: params["trade_status"],
		Sign:        sign,
	}
	err := cli.WithdrawCallback(req, func(r EPayCallbackReq) error {
		t.Logf("withdraw callback: tradeNo=%s status=%s", r.TradeNo, r.TradeStatus)
		return nil
	})
	if err != nil {
		t.Fatalf("withdraw callback failed: %v", err)
	}
}
