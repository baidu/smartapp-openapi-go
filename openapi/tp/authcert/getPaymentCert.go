package authcert

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetPaymentCertRequest 请求结构体
type GetPaymentCertRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
}

// 响应结构体

type GetPaymentCertResponsedata struct {
	URL string `json:"url"` // 认证页面调起 url
}

type GetPaymentCertResponse struct {
	Data      GetPaymentCertResponsedata `json:"data"`       // 响应参数
	Errno     int64                      `json:"errno"`      // 状态码
	ErrMsg    string                     `json:"msg"`        // 错误信息
	ErrorCode int64                      `json:"error_code"` // openapi 错误码
	ErrorMsg  string                     `json:"error_msg"`  // openapi 错误信息
}

// GetPaymentCert
func GetPaymentCert(params *GetPaymentCertRequest) (*GetPaymentCertResponsedata, error) {
	var (
		err        error
		defaultRet *GetPaymentCertResponsedata
	)
	respData := &GetPaymentCertResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/auth/paymentCertification")
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
	if respData.ErrorCode != 0 {
		return defaultRet, &utils.OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}

	if respData.Errno != 0 {
		return defaultRet, &utils.APIError{respData.Errno, respData.ErrMsg, respData}
	}
	return &respData.Data, nil
}
