package authprocess

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// RetrieveAuthCodeRequest 请求结构体
type RetrieveAuthCodeRequest struct {
	AccessToken string // 第三方平台的接口调用凭据
	AppID       int64  // 小程序 app_id
}

// 响应结构体

type RetrieveAuthCodeResponsedata struct {
	AuthorizationCode string `json:"authorization_code"` // 授权码
	ExpiresIn         int64  `json:"expires_in"`         // 授权码有效时间
}

type RetrieveAuthCodeResponse struct {
	Data      RetrieveAuthCodeResponsedata `json:"data"`       // 响应参数
	Errno     int64                        `json:"errno"`      // 状态码
	ErrMsg    string                       `json:"msg"`        // 错误信息
	ErrorCode int64                        `json:"error_code"` // openapi 错误码
	ErrorMsg  string                       `json:"error_msg"`  // openapi 错误信息
}

// RetrieveAuthCode
func RetrieveAuthCode(params *RetrieveAuthCodeRequest) (*RetrieveAuthCodeResponsedata, error) {
	var (
		err        error
		defaultRet *RetrieveAuthCodeResponsedata
	)
	respData := &RetrieveAuthCodeResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/auth/retrieve/authorizationcode")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("app_id", fmt.Sprintf("%v", params.AppID))

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
