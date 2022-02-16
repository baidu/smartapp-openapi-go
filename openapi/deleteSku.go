package openapi

import (
	"fmt"
)

// DeleteSkuRequest 请求结构体
type DeleteSkuRequest struct {
	AccessToken string // 接口调用凭证
	AppID       int64  // app_id
	PathList    string // 需要删除的资源 path 列表
}

// 响应结构体

type DeleteSkuResponse struct {
	Data      string `json:"data"`       // 响应参数
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// DeleteSku
func DeleteSku(params *DeleteSkuRequest) (string, error) {
	var (
		err        error
		defaultRet string
	)
	respData := &DeleteSkuResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/server/delete/sku")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("app_id", fmt.Sprintf("%v", params.AppID))
	client.AddPostParam("path_list", params.PathList)

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
