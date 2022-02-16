package openapi

import (
	"fmt"
)

// SubmitSitemapRequest 请求结构体
type SubmitSitemapRequest struct {
	AccessToken string // 小程序权限校验 Token
	AppID       int64  // app_id
	Desc        string // 描述信息
	Frequency   string // 更新频率 3-每天 4-每周
	Type        string // 类型 1-增量/更新； 0-下线/删除
	URL         string // sitemap 链接
}

// 响应结构体

type SubmitSitemapResponse struct {
	Data      bool   `json:"data"`       // 响应参数
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// SubmitSitemap
func SubmitSitemap(params *SubmitSitemapRequest) (bool, error) {
	var (
		err        error
		defaultRet bool
	)
	respData := &SubmitSitemapResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/access/submitsitemap")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("app_id", fmt.Sprintf("%v", params.AppID))
	client.AddPostParam("desc", params.Desc)
	client.AddPostParam("frequency", params.Frequency)
	client.AddPostParam("type", params.Type)
	client.AddPostParam("url", params.URL)

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

	return respData.Data, nil
}
