package template

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// AddToTemplateRequest 请求结构体
type AddToTemplateRequest struct {
	AccessToken string // 第三方平台的接口调用凭据
	DraftID     string // 草稿 id
	UserDesc    string // 自定义模板名称，30字以内
}

// 响应结构体

type AddToTemplateResponsedata struct {
	TemplateID int64 `json:"template_id"` // 返回模板 id
}

type AddToTemplateResponse struct {
	Data      AddToTemplateResponsedata `json:"data"`       // 响应参数
	Errno     int64                     `json:"errno"`      // 状态码
	ErrMsg    string                    `json:"msg"`        // 错误信息
	ErrorCode int64                     `json:"error_code"` // openapi 错误码
	ErrorMsg  string                    `json:"error_msg"`  // openapi 错误信息
}

// AddToTemplate
func AddToTemplate(params *AddToTemplateRequest) (*AddToTemplateResponsedata, error) {
	var (
		err        error
		defaultRet *AddToTemplateResponsedata
	)
	respData := &AddToTemplateResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/addtotemplate")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("draft_id", params.DraftID)
	client.AddPostParam("user_desc", params.UserDesc)

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
