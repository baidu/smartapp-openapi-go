package authprocess

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetTpTokenRequest 请求结构体
type GetTpTokenRequest struct {
	ClientID string // 第三方平台Key
	Ticket   string // 第三方平台服务器推送的 ticket，此 ticket 会定时推送，具体请见“1、 推送ticket协议”。
}

// 响应结构体

type GetTpTokenResponsedata struct {
	AccessToken string `json:"access_token"` // 第三方平台的接口调用凭据
	ExpiresIn   int64  `json:"expires_in"`   // 凭证有效时间（单位：秒）
	Scope       string `json:"scope"`        // 拥有的权限说明
}

type GetTpTokenResponse struct {
	Data      GetTpTokenResponsedata `json:"data"`       //
	Errno     int64                  `json:"errno"`      // 错误码；关于错误码的详细信息请参考 https://developer.baidu.com/wiki/index.php?title=docs/oauth/error 。
	ErrMsg    string                 `json:"msg"`        // 错误信息
	ErrorCode int64                  `json:"error_code"` // openapi 错误码
	ErrorMsg  string                 `json:"error_msg"`  // openapi 错误信息
}

// GetTpToken
func GetTpToken(params *GetTpTokenRequest) (*GetTpTokenResponsedata, error) {
	var (
		err        error
		defaultRet *GetTpTokenResponsedata
	)
	respData := &GetTpTokenResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/public/2.0/smartapp/auth/tp/token")
	client.AddGetParam("client_id", params.ClientID)
	client.AddGetParam("ticket", params.Ticket)
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
	if respData.ErrorCode != 0 {
		return defaultRet, &utils.OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}

	if respData.Errno != 0 {
		return defaultRet, &utils.APIError{respData.Errno, respData.ErrMsg, respData}
	}
	return &respData.Data, nil
}
