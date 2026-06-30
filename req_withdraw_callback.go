package go_epay

import (
	"encoding/json"
	"errors"

	"github.com/listenfengyang/go-epay/utils"
)

// WithdrawCallback 出金回调验签并处理
func (cli *Client) WithdrawCallback(req EPayCallbackReq, processor func(EPayCallbackReq) error) error {
	params := callbackToMap(req)
	if !utils.VerifyCallback(params, cli.Params.PayoutKey) {
		raw, _ := json.Marshal(req)
		cli.logger.Errorf("[EPay] withdraw callback verify failed: %s", string(raw))
		return errors.New("sign verify error")
	}
	return processor(req)
}
