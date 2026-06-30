package go_epay

import (
	"fmt"
	"sort"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-epay/utils"
	"github.com/mitchellh/mapstructure"
)

// Deposit 发起入金请求（Hosted Page 方式）
// POST https://api.epay365.biz/submit.php
func (cli *Client) Deposit(req EPayDepositReq) (*EPayDepositRsp, error) {
	req.PID = cli.Params.MerchantID
	if req.NotifyURL == "" {

		req.NotifyURL = cli.Params.NotifyURL
	}
	if req.ReturnURL == "" {

		req.ReturnURL = cli.Params.ReturnURL
	}

	var params map[string]string
	if err := mapstructure.Decode(req, &params); err != nil {
		return nil, fmt.Errorf("epay deposit: decode req failed: %w", err)
	}

	// 生成签名（原始值，不 URL encode）
	params["sign"] = utils.Sign(params, cli.Params.DepositKey)
	params["sign_type"] = "MD5"

	// 手动拼接原始 body（不 URL encode），与签名保持一致
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
	cli.logger.Infof("[EPay] deposit raw body: %s", rawBody)

	var result EPayDepositRsp
	rawURL := cli.Params.BaseURL + DepositPath

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
	cli.logger.Infof("[EPay] deposit response: %s", string(restLog))

	if err != nil {
		return nil, fmt.Errorf("epay deposit: request failed: %w", err)
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("epay deposit: http status %d body: %s", resp.StatusCode(), resp.Body())
	}
	if result.Code != 1 {
		return nil, fmt.Errorf("epay deposit: failed, code=%d msg=%s", result.Code, result.Msg)
	}
	return &result, nil
}
