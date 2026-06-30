package go_epay

import (
	"fmt"
	"sort"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-epay/utils"
	"github.com/mitchellh/mapstructure"
)

// DepositVia Server-to-Server 入金
func (cli *Client) DepositVia(req EPayDepositViaReq) (*EPayDepositViaRsp, error) {
	req.PID = cli.Params.MerchantID
	if req.NotifyURL == "" {
		req.NotifyURL = cli.Params.NotifyURL
	}
	if req.ReturnURL == "" {
		req.ReturnURL = cli.Params.ReturnURL
	}

	var params map[string]string
	if err := mapstructure.Decode(req, &params); err != nil {
		return nil, fmt.Errorf("epay deposit via: decode req failed: %w", err)
	}

	params["sign"] = utils.Sign(params, cli.Params.DepositKey)
	params["sign_type"] = "MD5"

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
	cli.logger.Infof("[EPay] deposit via raw body: %s", rawBody)

	var result EPayDepositViaRsp
	rawURL := cli.Params.BaseURL + DepositViaPath

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
	cli.logger.Infof("PSPResty#epay#deposit->%s", string(restLog))

	if err != nil {
		return nil, fmt.Errorf("epay deposit via: request failed: %w", err)
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("epay deposit via: http status %d body: %s", resp.StatusCode(), resp.Body())
	}
	if result.Code != 1 {
		return nil, fmt.Errorf("epay deposit via: failed, code=%d msg=%s", result.Code, result.Msg)
	}
	return &result, nil
}
