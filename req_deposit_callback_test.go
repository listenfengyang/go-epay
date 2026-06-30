package go_epay

import (
	"testing"

	"github.com/listenfengyang/go-epay/utils"
)

// TestDepositCallback_RealCallback 使用真实平台回调数据验证签名
// 回调原始 URL 参数:
// money=100.00&name=WangWei&out_trade_no=TV1782806065&param=&pid=Z100257
// &sign_type=MD5&trade_no=2026063015540000038351&trade_status=TRADE_SUCCESS
// &sign=1aecfce5b281e53649ce46ec90ba2cf5
func TestDepositCallback_RealCallback(t *testing.T) {
	cli := newTestClient()

	req := EPayCallbackReq{
		PID:         "Z100257",
		TradeNo:     "2026063015540000038351",
		OutTradeNo:  "TV1782806065",
		Name:        "WangWei",
		Money:       "100.00",
		TradeStatus: "TRADE_SUCCESS",
		Param:       "",
		Sign:        "1aecfce5b281e53649ce46ec90ba2cf5",
		SignType:    "MD5",
	}

	err := cli.DepositCallback(req, func(r EPayCallbackReq) error {
		t.Logf("callback processed: pid=%s tradeNo=%s outTradeNo=%s name=%s money=%s status=%s",
			r.PID, r.TradeNo, r.OutTradeNo, r.Name, r.Money, r.TradeStatus)
		return nil
	})
	if err != nil {
		t.Fatalf("callback failed: %v", err)
	}
}

// TestDepositCallback_MockSign 使用本地生成签名测试回调完整流程
func TestDepositCallback_MockSign(t *testing.T) {
	cli := newTestClient()

	params := map[string]string{
		"pid":          TestMerchantID,
		"trade_no":     "2026063015540000099999",
		"out_trade_no": "TV_TEST_ORDER_001",
		"name":         "WangWei",
		"money":        "100.00",
		"trade_status": "TRADE_SUCCESS",
	}
	sign := utils.Sign(params, DepositKey)
	t.Logf("generated sign: %s", sign)

	req := EPayCallbackReq{
		PID:         params["pid"],
		TradeNo:     params["trade_no"],
		OutTradeNo:  params["out_trade_no"],
		Name:        params["name"],
		Money:       params["money"],
		TradeStatus: params["trade_status"],
		Sign:        sign,
		SignType:    "MD5",
	}

	err := cli.DepositCallback(req, func(r EPayCallbackReq) error {
		t.Logf("callback processed: tradeNo=%s outTradeNo=%s status=%s", r.TradeNo, r.OutTradeNo, r.TradeStatus)
		return nil
	})
	if err != nil {
		t.Fatalf("callback failed: %v", err)
	}
}
