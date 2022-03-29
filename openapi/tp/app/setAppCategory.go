package app

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// SetAppCategoryRequest 请求结构体
type SetAppCategoryRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	Categorys   string // 类目类别
}

// 响应结构体

type SetAppCategoryResponsedata struct {
	ExamineID int64 `json:"examine_id"` //
}

type SetAppCategoryResponse struct {
	Data      SetAppCategoryResponsedata `json:"data"`       // 响应参数
	Errno     int64                      `json:"errno"`      // 状态码
	ErrMsg    string                     `json:"msg"`        // 错误信息
	ErrorCode int64                      `json:"error_code"` // openapi 错误码
	ErrorMsg  string                     `json:"error_msg"`  // openapi 错误信息
}

// SetAppCategory
func SetAppCategory(params *SetAppCategoryRequest) (*SetAppCategoryResponsedata, error) {
	var (
		err        error
		defaultRet *SetAppCategoryResponsedata
	)
	respData := &SetAppCategoryResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/app/category/update")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("categorys", params.Categorys)

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
