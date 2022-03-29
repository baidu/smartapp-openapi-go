package authprocess

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// RefreshOAuthTokenRequest 请求结构体
type RefreshOAuthTokenRequest struct {
	AccessToken  string // 第三方平台的接口调用凭据
	RefreshToken string // 接口调用凭据刷新令牌，有效期10年，使用后失效
	GrantType    string // 固定字符串： app_to_tp_refresh_token
}

// 响应结构体

type RefreshOAuthTokenResponse struct {
	AccessToken      string `json:"access_token"`      // 授权小程序的接口调用凭据
	Error            string `json:"error"`             // 错误码；关于错误码的详细信息请参考 https://developer.baidu.com/wiki/index.php?title=docs/oauth/error
	ErrorDescription string `json:"error_description"` // 错误描述信息，用来帮助理解和解决发生的错误
	ExpiresIn        int64  `json:"expires_in"`        // 小程序的Access Token的有效期，单位：秒，默认1小时
	RefreshToken     string `json:"refresh_token"`     // 接口调用凭据刷新令牌
	ErrorCode        int64  `json:"error_code"`        // openapi 错误码
	ErrorMsg         string `json:"error_msg"`         // openapi 错误信息
	Data             RefreshOAuthTokenResponsedata
}

type RefreshOAuthTokenResponsedata struct {
	AccessToken      string `json:"access_token"`      // 授权小程序的接口调用凭据
	Error            string `json:"error"`             // 错误码；关于错误码的详细信息请参考 https://developer.baidu.com/wiki/index.php?title=docs/oauth/error
	ErrorDescription string `json:"error_description"` // 错误描述信息，用来帮助理解和解决发生的错误
	ExpiresIn        int64  `json:"expires_in"`        // 小程序的Access Token的有效期，单位：秒，默认1小时
	RefreshToken     string `json:"refresh_token"`     // 接口调用凭据刷新令牌
}

// RefreshOAuthToken
func RefreshOAuthToken(params *RefreshOAuthTokenRequest) (*RefreshOAuthTokenResponsedata, error) {
	var (
		err        error
		defaultRet *RefreshOAuthTokenResponsedata
	)
	respData := &RefreshOAuthTokenResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/oauth/token")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("refresh_token", params.RefreshToken)
	client.AddGetParam("grant_type", params.GrantType)
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
	if respData.ErrorCode != 0 || respData.Error != "" {
		return defaultRet, &utils.OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}

	return &respData.Data, nil
}
