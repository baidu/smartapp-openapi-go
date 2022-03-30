package customermsg

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// SendCustomerMessageRequest 请求结构体
type SendCustomerMessageRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	UserType    int64  // 1:游客登录 2:百度账号登录
	OpenID      string // 用户的 OpenID
	MsgType     string // 消息类型 text:文本格式 image:图片链接
	Content     string // 文本消息内容，msg_type ="text" 时必填
	PicURL      string // 图片消息，msg_type ="image" 时必填
}

// 响应结构体

type SendCustomerMessageResponse struct {
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
	Data      SendCustomerMessageResponsedata
}

type SendCustomerMessageResponsedata struct {
	Errno int64  `json:"errno"` // 状态码
	Msg   string `json:"msg"`   // 状态描述
}

// SendCustomerMessage
func SendCustomerMessage(params *SendCustomerMessageRequest) (*SendCustomerMessageResponsedata, error) {
	var (
		err        error
		defaultRet *SendCustomerMessageResponsedata
	)
	respData := &SendCustomerMessageResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/message/custom/sendbytp")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("user_type", fmt.Sprintf("%v", params.UserType))
	client.AddPostParam("open_id", params.OpenID)
	client.AddPostParam("msg_type", params.MsgType)
	client.AddPostParam("content", params.Content)
	client.AddPostParam("pic_url", params.PicURL)

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
