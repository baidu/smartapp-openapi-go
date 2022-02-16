package openapi

import (
	"encoding/json"
)

// CheckPathRequest 请求结构体
type CheckPathRequest struct {
	AccessToken string   // 接口调用凭据
	Path        string   // 需要检测的页面 path(一次请求一个 path)，path 字符数不能超过 2460
	Type        []string // 检测策略，risk 为 url 里文字的内容违规检测，porn 为图片色情检测，ocr-word 为图片上文字的词表检测，ocr-lead 为图片上文字的诱导模型检测。可以多选，不传默认为 risk，参数值区分大小写
}

// 响应结构体

type CheckPathResponse struct {
	Data      string `json:"data"`       // 唯一标识 retrieveId，可用于精确查询检测结果、反馈误判
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// CheckPath
func CheckPath(params *CheckPathRequest) (string, error) {
	var (
		err        error
		defaultRet string
	)
	respData := &CheckPathResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeJSON).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/riskDetection/v2/asyncCheckPath")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	postData := map[string]interface{}{
		"path": params.Path,
		"type": params.Type,
	}
	bts, err := json.Marshal(postData)
	if err != nil {
		return defaultRet, err
	}
	client.SetBody(bts)

	err = client.Do()
	if err != nil {
		return defaultRet, err
	}
	err = client.Convert(respData)
	if err != nil {
		return defaultRet, err
	}
	if respData.ErrorCode != 0 {
		return defaultRet, &OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}
	if respData.Errno != 0 {
		return defaultRet, &APIError{respData.Errno, respData.ErrMsg, respData}
	}

	return respData.Data, nil
}
