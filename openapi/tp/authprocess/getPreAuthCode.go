package authprocess

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetPreAuthCodeRequest 请求结构体
type GetPreAuthCodeRequest struct {
	AccessToken string // 第三方平台的接口调用凭据
}

// 响应结构体

type GetPreAuthCodeResponsedata struct {
	ExpiresIn   int64  `json:"expires_in"`    // 凭证有效时间，单位：秒，默认20分钟
	PreAuthCode string `json:"pre_auth_code"` // 预授权码
}

type GetPreAuthCodeResponse struct {
	Data      GetPreAuthCodeResponsedata `json:"data"`       // 响应参数
	Errno     int64                      `json:"errno"`      // 状态码
	ErrMsg    string                     `json:"msg"`        // 错误信息
	ErrorCode int64                      `json:"error_code"` // openapi 错误码
	ErrorMsg  string                     `json:"error_msg"`  // openapi 错误信息
}

// GetPreAuthCode
func GetPreAuthCode(params *GetPreAuthCodeRequest) (*GetPreAuthCodeResponsedata, error) {
	var (
		err        error
		defaultRet *GetPreAuthCodeResponsedata
	)
	respData := &GetPreAuthCodeResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/tp/createpreauthcode")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)

	err = client.Do()
	if err != nil {
		return defaultRet, err
	}
	err = client.Convert(respData)
	if err != nil {
		return defaultRet, err
	}
	err = client.Convert(&respData.Data)
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
