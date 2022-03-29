package mobileauth

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// ApplyMobileAuthRequest 请求结构体
type ApplyMobileAuthRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	Reason      int64  // 申请原因（ 0："用于登录"；1 ："收货联系方式"；2 ："其他"）
	UsedScene   int64  // 使用场景（ 0："网络购物"1 ："账号下信息内容同步"；2 ："票务预订"；3 ："业务办理"；4 ："信息查询（如社保、公积金查询"；5 ：预约"）
	SceneDesc   string // 使用场景描述
	SceneDemo   string // 使用场景 demo （场景实例图片）
}

// 响应结构体

type ApplyMobileAuthResponsedata struct {
}

type ApplyMobileAuthResponse struct {
	Data      ApplyMobileAuthResponsedata `json:"data"`       // 响应参数
	Errno     int64                       `json:"errno"`      // 状态码
	ErrMsg    string                      `json:"msg"`        // 错误信息
	ErrorCode int64                       `json:"error_code"` // openapi 错误码
	ErrorMsg  string                      `json:"error_msg"`  // openapi 错误信息
}

// ApplyMobileAuth
func ApplyMobileAuth(params *ApplyMobileAuthRequest) (*ApplyMobileAuthResponsedata, error) {
	var (
		err        error
		defaultRet *ApplyMobileAuthResponsedata
	)
	respData := &ApplyMobileAuthResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/app/apply/mobileauth")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("reason", fmt.Sprintf("%v", params.Reason))
	client.AddPostParam("used_scene", fmt.Sprintf("%v", params.UsedScene))
	client.AddPostParam("scene_desc", params.SceneDesc)
	client.AddPostParam("scene_demo", params.SceneDemo)

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
