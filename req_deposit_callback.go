package go_epay

import (
	"encoding/json"
	"errors"

	"github.com/listenfengyang/go-epay/utils"
)

// DepositCallback 入金回调验签并处理
// 回调成功后需返回字符串 "success"
func (cli *Client) DepositCallback(req EPayCallbackReq, processor func(EPayCallbackReq) error) error {
	params := callbackToMap(req)
	if !utils.VerifyCallback(params, cli.Params.DepositKey) {
		raw, _ := json.Marshal(req)
		cli.logger.Errorf("[EPay] deposit callback verify failed: %s", string(raw))
		return errors.New("sign verify error")
	}
	return processor(req)
}

// callbackToMap 将回调结构体转为 map[string]interface{}，key 使用 json tag
func callbackToMap(req EPayCallbackReq) map[string]interface{} {
	return map[string]interface{}{
		"pid":          req.PID,
		"trade_no":     req.TradeNo,
		"out_trade_no": req.OutTradeNo,
		"type":         req.Type,
		"name":         req.Name,
		"money":        req.Money,
		"trade_status": req.TradeStatus,
		"param":        req.Param,
		"sign":         req.Sign,
		"sign_type":    req.SignType,
	}
}
