package utils

import "testing"

func TestSign(t *testing.T) {
	params := map[string]string{
		"pid":          "1001",
		"type":         "alipay",
		"out_trade_no": "ORDER123",
		"notify_url":   "https://example.com/notify",
		"name":         "test",
		"money":        "100.00",
		"clientip":     "1.2.3.4",
	}
	sign := Sign(params, "testkey")
	if sign == "" {
		t.Fatal("sign is empty")
	}
	t.Logf("sign: %s", sign)

	params["sign"] = sign
	if !Verify(params, "testkey") {
		t.Fatal("verify failed")
	}
}
