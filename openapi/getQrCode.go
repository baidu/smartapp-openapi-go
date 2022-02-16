package openapi

import (
	"fmt"
)

// GetQrCodeRequest 请求结构体
type GetQrCodeRequest struct {
	AccessToken string // 接口调用凭证
	Path        string // 扫码进入的小程序页面路径，最大长度 4000 字节，可以为空
	Width       int64  // 二维码的宽度（单位：px）。最小 280px，最大 1280px
	Mf          int64  // 是否包含二维码内嵌 logo 标识，1001 为不包含，默认包含
	ImgFlag     int64  // 返回值选项，默认或传 1 时只返回二维码 base64 编码字符串，传 0 只返回 url
}

// 响应结构体

type GetQrCodeResponsedata struct {
	Base64Str string `json:"base64_str"` // 二维码 base64 编码字符串
	URL       string `json:"url"`        // 小程序二维码短链
}

type GetQrCodeResponse struct {
	Data      GetQrCodeResponsedata `json:"data"`       // 响应对象
	ErrMsg    string                `json:"msg"`        // 错误信息
	Errno     int64                 `json:"errno"`      // 错误码
	RequestID string                `json:"request_id"` // 请求 ID，标识一次请求
	Timestamp int64                 `json:"timestamp"`  // 时间戳
	ErrorCode int64                 `json:"error_code"` // openapi 错误码
	ErrorMsg  string                `json:"error_msg"`  // openapi 错误信息
}

// GetQrCode
func GetQrCode(params *GetQrCodeRequest) (*GetQrCodeResponsedata, error) {
	var (
		err        error
		defaultRet *GetQrCodeResponsedata
	)
	respData := &GetQrCodeResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/qrcode/getv2")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("path", params.Path)
	client.AddPostParam("width", fmt.Sprintf("%v", params.Width))
	client.AddPostParam("mf", fmt.Sprintf("%v", params.Mf))
	client.AddPostParam("img_flag", fmt.Sprintf("%v", params.ImgFlag))

	err = client.Do()
	if err != nil {
		return defaultRet, err
	}
	err = client.Convert(respData)
	if err != nil {
		return defaultRet, err
	}
	if respData.ErrorCode != 0 {
		return defaultRet, &OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}
	if respData.Errno != 0 {
		return defaultRet, &APIError{respData.Errno, respData.ErrMsg, respData}
	}

	return &respData.Data, nil
}
