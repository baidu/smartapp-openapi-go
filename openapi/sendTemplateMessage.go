package openapi

import (
	"fmt"

)

// SendTemplateMessageRequest 请求结构体
type SendTemplateMessageRequest struct {
	AccessToken  string // 接口调用凭证
	TemplateID   string // 小程序模板 ID
	TouserOpenID string // 接收者 open_id 参数不能为空。open_id 为百度用户登录唯一标识，可以通过 SessionKey 获得
	Data         string // 发送消息内容
	Page         string // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数，（示例 index?foo=bar），该字段不填默认跳转至首页
	SceneID      string // 场景 id ，例如 formId、orderId、payId。formId 为页面内 form 组件的report-submit属性为 true 时返回 formId ，详见 form 表单
	SceneType    int64  // 场景 type 。1：表单；2：百度收银台订单；3：直连订单
}

// 响应结构体

type SendTemplateMessageResponsedata struct {
	MsgKey int64 `json:"msg_key"` // 消息id
}

type SendTemplateMessageResponse struct {
	Data      SendTemplateMessageResponsedata `json:"data"`       // 响应对象
	Errno     int64                           `json:"errno"`      // 错误码
	ErrMsg    string                          `json:"msg"`        // 错误信息
	ErrorCode int64                           `json:"error_code"` // openapi 错误码
	ErrorMsg  string                          `json:"error_msg"`  // openapi 错误信息
}

// SendTemplateMessage
func SendTemplateMessage(params *SendTemplateMessageRequest) (*SendTemplateMessageResponsedata, error) {
	var (
		err        error
		defaultRet *SendTemplateMessageResponsedata
	)
	respData := &SendTemplateMessageResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/send")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("template_id", params.TemplateID)
	client.AddPostParam("touser_openId", params.TouserOpenID)
	client.AddPostParam("data", params.Data)
	client.AddPostParam("page", params.Page)
	client.AddPostParam("scene_id", params.SceneID)
	client.AddPostParam("scene_type", fmt.Sprintf("%v", params.SceneType))

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
