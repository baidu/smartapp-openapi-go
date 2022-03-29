package search

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// SubmitSitemapByAPIRequest 请求结构体
type SubmitSitemapByAPIRequest struct {
	AccessToken string // 第三方平台的接口调用凭据
	Type        int64  // 上传级别 0：周级别，一周左右生效；1：天级别，2~3天生效
	URLList     string // url集合；上传级别上限，0：每天3000条，1：每天5000条 多个,分割
}

// 响应结构体

type SubmitSitemapByAPIResponsedata struct {
}

type SubmitSitemapByAPIResponse struct {
	Data      SubmitSitemapByAPIResponsedata `json:"data"`       // 响应参数
	Errno     int64                          `json:"errno"`      // 状态码
	ErrMsg    string                         `json:"msg"`        // 错误信息
	ErrorCode int64                          `json:"error_code"` // openapi 错误码
	ErrorMsg  string                         `json:"error_msg"`  // openapi 错误信息
}

// SubmitSitemapByAPI
func SubmitSitemapByAPI(params *SubmitSitemapByAPIRequest) (*SubmitSitemapByAPIResponsedata, error) {
	var (
		err        error
		defaultRet *SubmitSitemapByAPIResponsedata
	)
	respData := &SubmitSitemapByAPIResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/access/submit/sitemap")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("type", fmt.Sprintf("%v", params.Type))
	client.AddPostParam("url_list", params.URLList)

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
