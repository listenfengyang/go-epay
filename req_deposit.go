package go_epay

import (
	"fmt"
	"html"
	"sort"
	"strings"

	"github.com/listenfengyang/go-epay/utils"
	"github.com/mitchellh/mapstructure"
)

// Deposit 发起入金请求（Hosted Page 方式）
//
// /submit.php 只接受 POST 表单，不支持 GET 请求。
// 本方法不发起 HTTP 请求，本地生成完整 HTML 页面：
//   - PayURL:   表单 action 地址（https://api.epay365.biz/submit.php）
//   - FormHTML: 完整的自动提交 HTML 页面，后端直接以 Content-Type: text/html 返回给浏览器，
//     浏览器渲染后立即 POST 跳转到支付页
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

	// 生成签名
	params["sign"] = utils.Sign(params, cli.Params.DepositKey)
	params["sign_type"] = "MD5"

	// 按 key 排序
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	actionURL := cli.Params.BaseURL + cli.Params.DepositPath

	// 构造完整的自动提交 HTML 页面
	var sb strings.Builder
	sb.WriteString(`<!DOCTYPE html><html><head><meta charset="utf-8"><title>Redirecting...</title></head><body>`)
	sb.WriteString(`<form id="epay-deposit-form" method="POST" action="`)
	sb.WriteString(html.EscapeString(actionURL))
	sb.WriteString(`">`)
	for _, k := range keys {
		sb.WriteString(`<input type="hidden" name="`)
		sb.WriteString(html.EscapeString(k))
		sb.WriteString(`" value="`)
		sb.WriteString(html.EscapeString(params[k]))
		sb.WriteString(`">`)
	}
	sb.WriteString(`</form>`)
	sb.WriteString(`<script>document.getElementById("epay-deposit-form").submit();</script>`)
	sb.WriteString(`</body></html>`)
	formHTML := sb.String()

	cli.logger.Infof("[EPay] deposit actionURL: %s", actionURL)

	return &EPayDepositRsp{
		Code:     1,
		PayURL:   actionURL,
		FormHTML: formHTML,
	}, nil
}
