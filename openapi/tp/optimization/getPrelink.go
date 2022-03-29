package optimization

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetPrelinkRequest 请求结构体
type GetPrelinkRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
}

// 响应结构体

type GetPrelinkResponse struct {
	Data      []string `json:"data"`       // 响应参数
	Errno     int64    `json:"errno"`      // 状态码
	ErrMsg    string   `json:"msg"`        // 错误信息
	ErrorCode int64    `json:"error_code"` // openapi 错误码
	ErrorMsg  string   `json:"error_msg"`  // openapi 错误信息
}

// GetPrelink
func GetPrelink(params *GetPrelinkRequest) ([]string, error) {
	var (
		err        error
		defaultRet []string
	)
	respData := &GetPrelinkResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/prelink/get")
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
	return respData.Data, nil
}
