package msgtemplate

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// AddMsgTemplateRequest 请求结构体
type AddMsgTemplateRequest struct {
	AccessToken   string // 授权小程序的接口调用凭据
	ID            string // 模板标题 id
	KeywordIDList string // 模板关键词 id 列表，如[1,2,3]；
}

// 响应结构体

type AddMsgTemplateResponsedata struct {
	TemplateID string `json:"template_id"` // 添加至帐号下的模板id，发送小程序模板消息时所需
}

type AddMsgTemplateResponse struct {
	Data      AddMsgTemplateResponsedata `json:"data"`       // 响应参数
	Errno     int64                      `json:"errno"`      // 状态码
	ErrMsg    string                     `json:"msg"`        // 错误信息
	ErrorCode int64                      `json:"error_code"` // openapi 错误码
	ErrorMsg  string                     `json:"error_msg"`  // openapi 错误信息
}

// AddMsgTemplate
func AddMsgTemplate(params *AddMsgTemplateRequest) (*AddMsgTemplateResponsedata, error) {
	var (
		err        error
		defaultRet *AddMsgTemplateResponsedata
	)
	respData := &AddMsgTemplateResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/add")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("id", params.ID)
	client.AddPostParam("keyword_id_list", params.KeywordIDList)

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
