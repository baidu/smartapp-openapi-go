package search

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// SaveURLMappingRequest 请求结构体
type SaveURLMappingRequest struct {
	AccessToken  string // 授权小程序的接口调用凭据
	Content      string // json字符串
	CreateMethod string // 每页数量，最大100
}

// 响应结构体

type SaveURLMappingResponse struct {
	Data      bool   `json:"data"`       // 响应参数
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// SaveURLMapping
func SaveURLMapping(params *SaveURLMappingRequest) (bool, error) {
	var (
		err        error
		defaultRet bool
	)
	respData := &SaveURLMappingResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/flow/saveurlmapping")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("content", params.Content)
	client.AddPostParam("create_method", params.CreateMethod)

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
	return respData.Data, nil
}
