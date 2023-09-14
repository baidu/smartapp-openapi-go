package openapi

// SubscribeSendRequest 请求结构体
type SubscribeSendRequest struct {
	AccessToken  string // 接口调用凭证
	TemplateID   string // 所需下发的模板消息的id
	TouserOpenID string // 接收者open_id
	SubscribeID  string // 订阅 Id ，发送订阅类模板消息时所使用的唯一标识符，开发者自定义的subscribe-id 字段。注意：同一用户在同一个订阅id 下的多次授权不累积下发权限，只能下发一条。若要订阅多条，需要不同订阅 id
	Data         string // 消息内容
	Page         string // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数，示例index?foo=bar，该字段不填默认跳转至首页
}

// 响应结构体

type SubscribeSendResponsedata struct {
	MsgKey int64 `json:"msg_key"` // 消息 id
}

type SubscribeSendResponse struct {
	Data      SubscribeSendResponsedata `json:"data"`       // 响应对象
	Errno     int64                     `json:"errno"`      // 错误码
	ErrMsg    string                    `json:"msg"`        // 错误信息
	ErrorCode int64                     `json:"error_code"` // openapi 错误码
	ErrorMsg  string                    `json:"error_msg"`  // openapi 错误信息
}

// SubscribeSend
func SubscribeSend(params *SubscribeSendRequest) (*SubscribeSendResponsedata, error) {
	var (
		err        error
		defaultRet *SubscribeSendResponsedata
	)
	respData := &SubscribeSendResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/message/subscribe/send")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("template_id", params.TemplateID)
	client.AddPostParam("touser_openId", params.TouserOpenID)
	client.AddPostParam("subscribe_id", params.SubscribeID)
	client.AddPostParam("data", params.Data)
	client.AddPostParam("page", params.Page)

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
