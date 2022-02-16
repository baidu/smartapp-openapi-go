package openapi

import (
	"strconv"
)

// ResetComponentStorageRequest 请求结构体
type ResetComponentStorageRequest struct {
	AccessToken string // 接口调用凭证
}

// 响应结构体
type ResetComponentStorageResponsedata struct {
	ErrMsg string
	Errno  string
	Logid  string
}

type ResetComponentStorageResponse struct {
	ErrMsg    string `json:"msg"`        // 错误信息
	Errno     string `json:"errno"`      // 错误码
	ErrorCode int64  `json:"error_code"` // 错误码
	ErrorMsg  string `json:"error_msg"`  // 错误信息
	Logid     string `json:"logid"`      // 请求 ID，标识一次请求
}

// ResetComponentStorage
func ResetComponentStorage(params *ResetComponentStorageRequest) (*ResetComponentStorageResponsedata, error) {
	var (
		err        error
		defaultRet *ResetComponentStorageResponsedata
	)
	respData := &ResetComponentStorageResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/storage/component/reset")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)

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
	errno, err := strconv.ParseInt(respData.Errno, 10, 64)
	if err != nil {
		return defaultRet, err
	}
	if errno != 0 {
		return defaultRet, &APIError{errno, respData.ErrMsg, respData}
	}

	resData := &ResetComponentStorageResponsedata{
		Errno:  respData.Errno,
		ErrMsg: respData.ErrMsg,
		Logid:  respData.Logid,
	}

	return resData, nil
}
