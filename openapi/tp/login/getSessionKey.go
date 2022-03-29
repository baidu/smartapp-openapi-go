package login

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetSessionKeyRequest 请求结构体
type GetSessionKeyRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	Code        string // 调用 swan.login 后获取的 code
	GrantType   string // 授权类型，固定字符串：“authorization_code”
}

// 响应结构体

type GetSessionKeyResponse struct {
	Errno            int64  `json:"errno"`             // 错误码
	ErrorDescription string `json:"error_description"` // 错误描述信息，用来帮助理解和解决发生的错误
	ErrMsg           string `json:"msg"`               // 错误信息
	ErrorCode        int64  `json:"error_code"`        // openapi 错误码
	ErrorMsg         string `json:"error_msg"`         // openapi 错误信息
	Data             GetSessionKeyResponsedata
}

type GetSessionKeyResponsedata struct {
	Errno            int64  `json:"errno"`             // 错误码
	ErrorDescription string `json:"error_description"` // 错误描述信息，用来帮助理解和解决发生的错误
	Msg              string `json:"msg"`               // 错误描述
}

// GetSessionKey
func GetSessionKey(params *GetSessionKeyRequest) (*GetSessionKeyResponsedata, error) {
	var (
		err        error
		defaultRet *GetSessionKeyResponsedata
	)
	respData := &GetSessionKeyResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/oauth/getsessionkeybycode")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("code", params.Code)
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
	if respData.ErrorCode != 0 {
		return defaultRet, &utils.OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}

	return &respData.Data, nil
}
