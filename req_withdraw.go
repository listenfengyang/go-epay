package go_epay

import (
	"fmt"
	"sort"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-epay/utils"
	"github.com/mitchellh/mapstructure"
)

// Withdraw 发起出金请求
// POST https://api.epay365.biz/payout.php
func (cli *Client) Withdraw(req EPayWithdrawReq) (*EPayWithdrawRsp, error) {
	// 补充固定参数
	req.PID = cli.Params.MerchantID
	if req.NotifyURL == "" {
		req.NotifyURL = cli.Params.WithdrawNotifyURL
	}

	// struct -> map[string]string
	var params map[string]string
	if err := mapstructure.Decode(req, &params); err != nil {
		return nil, fmt.Errorf("epay withdraw: decode req failed: %w", err)
	}

	// 生成签名
	params["sign"] = utils.Sign(params, cli.Params.PayoutKey)
	params["sign_type"] = "MD5"

	// 构造有序 form body（便于日志排查）
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, params[k]))
	}
	rawBody := strings.Join(parts, "&")
	cli.logger.Infof("[EPay] withdraw raw body: %s", rawBody)

	var result EPayWithdrawRsp
	rawURL := cli.Params.BaseURL + WithdrawPath

	resp, err := cli.ryClient.
		R().
		SetBody(strings.NewReader(rawBody)).
		SetFormData(params).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp))
	cli.logger.Infof("PSPResty#epay#withdraw->%s", string(restLog))

	if err != nil {
		return nil, fmt.Errorf("epay withdraw: request failed: %w", err)
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("epay withdraw: http status %d body: %s", resp.StatusCode(), resp.Body())
	}
	if result.Code != 1 {
		return nil, fmt.Errorf("epay withdraw: failed, code=%d msg=%s", result.Code, result.Msg)
	}
	return &result, nil
}
