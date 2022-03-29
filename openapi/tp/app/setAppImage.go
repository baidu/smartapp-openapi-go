package app

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// SetAppImageRequest 请求结构体
type SetAppImageRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	ImageURL    string // 小程序icon地址，可以通过图片上传接口获取icon的url。
}

// 响应结构体

type SetAppImageResponsedata struct {
}

type SetAppImageResponse struct {
	Data      SetAppImageResponsedata `json:"data"`       // 响应参数
	Errno     int64                   `json:"errno"`      // 状态码
	ErrMsg    string                  `json:"msg"`        // 错误信息
	ErrorCode int64                   `json:"error_code"` // openapi 错误码
	ErrorMsg  string                  `json:"error_msg"`  // openapi 错误信息
}

// SetAppImage
func SetAppImage(params *SetAppImageRequest) (*SetAppImageResponsedata, error) {
	var (
		err        error
		defaultRet *SetAppImageResponsedata
	)
	respData := &SetAppImageResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/app/modifyheadimage")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("image_url", params.ImageURL)

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
