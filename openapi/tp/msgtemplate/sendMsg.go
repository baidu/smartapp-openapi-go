package msgtemplate

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// SendMsgRequest 请求结构体
type SendMsgRequest struct {
	AccessToken  string      // 授权小程序的接口调用凭据
	TemplateID   string      // 模板 id ，发送小程序模板消息时所需
	TouserOpenID string      // 接收者 open_id
	Data         string      // {"keyword1": {"value": "2018-09-06"},"keyword2": {"value": "kfc"}}
	Page         interface{} // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数，（示例index?foo=bar），该字段不填则模板无跳转。
	SceneID      string      // 场景 id，例如表单 id 和订单 id
	SceneType    int64       // 场景type，1：表单；2：百度收银台订单；3:直连订单
	Title        string      // 标题
	Ext          interface{} // {"xzh_id":111,"category_id":15}
}

// 响应结构体

type SendMsgResponsedata struct {
	MsgKey int64 `json:"msg_key"` // 消息 id
}

type SendMsgResponse struct {
	Data      SendMsgResponsedata `json:"data"`       // 响应参数
	Errno     int64               `json:"errno"`      // 状态码
	ErrMsg    string              `json:"msg"`        // 错误信息
	ErrorCode int64               `json:"error_code"` // openapi 错误码
	ErrorMsg  string              `json:"error_msg"`  // openapi 错误信息
}

// SendMsg
func SendMsg(params *SendMsgRequest) (*SendMsgResponsedata, error) {
	var (
		err        error
		defaultRet *SendMsgResponsedata
	)
	respData := &SendMsgResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/sendmessage")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("template_id", params.TemplateID)
	client.AddPostParam("touser_openId", params.TouserOpenID)
	client.AddPostParam("data", params.Data)
	client.AddPostParam("page", params.Page)
	client.AddPostParam("scene_id", params.SceneID)
	client.AddPostParam("scene_type", fmt.Sprintf("%v", params.SceneType))
	client.AddPostParam("title", params.Title)
	client.AddPostParam("ext", params.Ext)

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
