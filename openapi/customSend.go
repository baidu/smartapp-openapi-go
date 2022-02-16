package openapi

import (
	"fmt"
)

// CustomSendRequest 请求结构体
type CustomSendRequest struct {
	AccessToken string // 接口调用凭证
	UserType    int64  // 1:游客登录 2:百度账号登录
	OpenID      string // 用户的 OpenID
	MsgType     string // 消息类型 text:文本格式 image:图片链接
	Content     string // 文本消息内容，msg_type ="text" 时必填
	PicURL      string // 图片消息，msg_type ="image" 时必填
}

// 响应结构体
// 响应结构体
type CustomSendResponsedata struct {
	Errno int64
	Msg   string
}

type CustomSendResponse struct {
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// CustomSend
func CustomSend(params *CustomSendRequest) (*CustomSendResponsedata, error) {
	var (
		err        error
		defaultRet *CustomSendResponsedata
	)
	respData := &CustomSendResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/message/custom/send")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
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
		return defaultRet, &OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}
	if respData.Errno != 0 {
		return defaultRet, &APIError{respData.Errno, respData.ErrMsg, respData}
	}

	resData := &CustomSendResponsedata{
		Errno: respData.Errno,
		Msg:   respData.ErrMsg,
	}

	return resData, nil
}
