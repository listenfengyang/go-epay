package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// Sign 生成 epay 签名
// 规则：
//  1. 过滤 sign/sign_type 及空值
//  2. 按参数名 ASCII 升序排列
//  3. 拼接为 a=b&c=d 格式（不 URL encode）
//  4. 末尾追加商户密钥 KEY（直接拼接，无分隔符）
//  5. MD5 整个字符串，结果小写
func Sign(params map[string]string, key string) string {
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if k == "sign" || k == "sign_type" || v == "" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, params[k]))
	}
	raw := strings.Join(parts, "&") + key

	fmt.Printf("[EPay Sign] rawString: %s\n", raw)

	hash := md5.Sum([]byte(raw))
	return hex.EncodeToString(hash[:])
}

// Verify 验证签名（map[string]string）
func Verify(params map[string]string, key string) bool {
	sign, ok := params["sign"]
	if !ok || sign == "" {
		return false
	}
	return Sign(params, key) == sign
}

// VerifyCallback 验证回调签名（map[string]interface{}）
func VerifyCallback(params map[string]interface{}, key string) bool {
	sign, ok := params["sign"]
	if !ok {
		return false
	}
	signStr, ok := sign.(string)
	if !ok || signStr == "" {
		return false
	}
	filtered := make(map[string]string, len(params))
	for k, v := range params {
		if k == "sign" || k == "sign_type" {
			continue
		}
		if s, ok := v.(string); ok && s != "" {
			filtered[k] = s
		}
	}
	return Sign(filtered, key) == signStr
}
