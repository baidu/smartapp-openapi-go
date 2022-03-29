package app

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// ControlAppFlowRequest 请求结构体
type ControlAppFlowRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	Type        int64  // 流量下线开关状态，1：开启流量，2：下线流量
}

// 响应结构体

type ControlAppFlowResponse struct {
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
	Data      ControlAppFlowResponsedata
}

type ControlAppFlowResponsedata struct {
	Errno int64  `json:"errno"` // 状态码
	Msg   string `json:"msg"`   // 状态描述
}

// ControlAppFlow
func ControlAppFlow(params *ControlAppFlowRequest) (*ControlAppFlowResponsedata, error) {
	var (
		err        error
		defaultRet *ControlAppFlowResponsedata
	)
	respData := &ControlAppFlowResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/appflow/control")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("type", fmt.Sprintf("%v", params.Type))

	err = client.Do()
	if err != nil {
		return defaultRet, err
	}
	err = client.Convert(respData)
	if err != nil {
		return defaultRet, err
	}
	if respData.ErrorCode != 0 {
		return defaultRet, &utils.OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}

	if respData.Errno != 0 {
		return defaultRet, &utils.APIError{respData.Errno, respData.ErrMsg, respData}
	}
	return &respData.Data, nil
}
