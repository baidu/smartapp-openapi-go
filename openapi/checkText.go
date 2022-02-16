package openapi

import (
	"encoding/json"
)

// CheckTextRequest 请求结构体
type CheckTextRequest struct {
	AccessToken string   // 接口调用凭据
	Content     string   // 检测文本，文本的字节数不能超过 10KB
	Type        []string // 检测策略，risk 为内容违规检测，lead 为诱导违规检测。可以多选，不传默认为 risk，参数值区分大小写
}

// 响应结构体

type CheckTextResponsedataresItem struct {
	Errno int64  `json:"errno"` // 错误码
	Msg   string `json:"msg"`   // 错误信息
	Type  string `json:"type"`  // 检测策略
}

type CheckTextResponsedata struct {
	Res        []CheckTextResponsedataresItem `json:"res"`        // 是一个对象数组，返回每一种检测策略的结果
	RetrieveID string                         `json:"retrieveId"` // 调用误判反馈接口时需要该返回值
}

type CheckTextResponse struct {
	Data      CheckTextResponsedata `json:"data"`       // 响应对象
	Errno     int64                 `json:"errno"`      // 错误码
	ErrMsg    string                `json:"msg"`        // 错误信息
	ErrorCode int64                 `json:"error_code"` // openapi 错误码
	ErrorMsg  string                `json:"error_msg"`  // openapi 错误信息
}

// CheckText
func CheckText(params *CheckTextRequest) (*CheckTextResponsedata, error) {
	var (
		err        error
		defaultRet *CheckTextResponsedata
	)
	respData := &CheckTextResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeJSON).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/riskDetection/v2/syncCheckText")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	postData := map[string]interface{}{
		"content": params.Content,
		"type":    params.Type,
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

	return &respData.Data, nil
}
