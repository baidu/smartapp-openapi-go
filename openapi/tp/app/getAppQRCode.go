package app

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAppQRCodeRequest 请求结构体
type GetAppQRCodeRequest struct {
	AccessToken string      // 授权小程序的接口调用凭据
	Path        interface{} // 自定义打开路径
	PackageID   interface{} // 可指定代码包id(只支持审核、开发、线上版本)，不传默认线上版本
	Width       interface{} // 默认 200px ，最大 1280px ，最小 200px
}

// 响应结构体

type GetAppQRCodeResponse struct {
	ContentType string
	Data        GetAppQRCodeResponsedata
	Errno       int64  `json:"errno"`      // 状态码
	ErrMsg      string `json:"msg"`        // 错误信息
	ErrorCode   int64  `json:"error_code"` // openapi 错误码
	ErrorMsg    string `json:"error_msg"`  // openapi 错误信息
}

type GetAppQRCodeResponsedata []byte

// GetAppQRCode
func GetAppQRCode(params *GetAppQRCodeRequest) (*GetAppQRCodeResponse, error) {
	var (
		err        error
		defaultRet *GetAppQRCodeResponse
	)
	respData := &GetAppQRCodeResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypePNG).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/app/qrcode")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("path", params.Path)
	client.AddGetParam("package_id", params.PackageID)
	client.AddGetParam("width", params.Width)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)

	err = client.Do()
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
	respData.ContentType = utils.ConverterTypePNG
	return respData, nil
}
