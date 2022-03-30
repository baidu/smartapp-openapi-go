package authcert

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetFaceAuthRequest 请求结构体
type GetFaceAuthRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
}

// 响应结构体

type GetFaceAuthResponsedata struct {
	LegalFaceStatus string `json:"legal_face_status"` // 认证任务状态(1：初始化；2：通过；4：拒绝；6：进行中)
	Qrcode          string `json:"qrcode"`            // 二维码(base64编码图片)
	Timeout         string `json:"timeout"`           // 过期时间
}

type GetFaceAuthResponse struct {
	Data      GetFaceAuthResponsedata `json:"data"`       // 响应参数
	Errno     int64                   `json:"errno"`      // 状态码
	ErrMsg    string                  `json:"msg"`        // 错误信息
	ErrorCode int64                   `json:"error_code"` // openapi 错误码
	ErrorMsg  string                  `json:"error_msg"`  // openapi 错误信息
}

// GetFaceAuth
func GetFaceAuth(params *GetFaceAuthRequest) (*GetFaceAuthResponsedata, error) {
	var (
		err        error
		defaultRet *GetFaceAuthResponsedata
	)
	respData := &GetFaceAuthResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/auth/faceauthen")
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
