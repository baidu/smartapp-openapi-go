package template

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// DelTemplateRequest 请求结构体
type DelTemplateRequest struct {
	AccessToken string // 第三方平台的接口调用凭据
	TemplateID  string // 模板 id
}

// 响应结构体

type DelTemplateResponse struct {
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
	Data      DelTemplateResponsedata
}

type DelTemplateResponsedata struct {
	Errno int64  `json:"errno"` // 状态码
	Msg   string `json:"msg"`   // 状态描述
}

// DelTemplate
func DelTemplate(params *DelTemplateRequest) (*DelTemplateResponsedata, error) {
	var (
		err        error
		defaultRet *DelTemplateResponsedata
	)
	respData := &DelTemplateResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/deltemplate")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("template_id", params.TemplateID)

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
