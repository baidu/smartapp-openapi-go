package login

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetUnionIDRequest 请求结构体
type GetUnionIDRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	OpenID      string // 用户 openid ，需要经过用户登录授权
}

// 响应结构体

type GetUnionIDResponsedata struct {
	UnionID string `json:"union_id"` // 对应的 union id
}

type GetUnionIDResponse struct {
	Data      GetUnionIDResponsedata `json:"data"`       // 详细数据，errno 为 0 的情况下才有意义
	Errno     int64                  `json:"errno"`      // 状态码
	ErrMsg    string                 `json:"msg"`        // 错误信息
	ErrorCode int64                  `json:"error_code"` // openapi 错误码
	ErrorMsg  string                 `json:"error_msg"`  // openapi 错误信息
}

// GetUnionID
func GetUnionID(params *GetUnionIDRequest) (*GetUnionIDResponsedata, error) {
	var (
		err        error
		defaultRet *GetUnionIDResponsedata
	)
	respData := &GetUnionIDResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/unionId/get")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("open_id", params.OpenID)
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
