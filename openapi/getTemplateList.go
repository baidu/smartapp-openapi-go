package openapi

import (
	"fmt"
)

// GetTemplateListRequest 请求结构体
type GetTemplateListRequest struct {
	AccessToken string // 接口调用凭证
	Offset      int64  // 用于分页，表示从 offset 开始，默认值为 0
	Count       int64  // 用于分页，表示拉取 count 条记录，默认值为 0 ，最大为 20
}

// 响应结构体

type GetTemplateListResponsedatalistItem struct {
	Content    string `json:"content"`     // 模板内容
	Example    string `json:"example"`     // 模板内容示例
	TemplateID string `json:"template_id"` // 模板 id ，发送小程序模板消息时所需
	Title      string `json:"title"`       // 模板标题
}

type GetTemplateListResponsedata struct {
	List       []GetTemplateListResponsedatalistItem `json:"list"`        // 模板对象数组，对象中包含模板 id 、模板标题 title 、模板内容 content 、模板内容示例 example
	TotalCount int64                                 `json:"total_count"` // 模板库标题总数
}

type GetTemplateListResponse struct {
	Data      GetTemplateListResponsedata `json:"data"`       // 响应对象
	Errno     int64                       `json:"errno"`      // 错误码
	ErrMsg    string                      `json:"msg"`        // 错误信息
	ErrorCode int64                       `json:"error_code"` // openapi 错误码
	ErrorMsg  string                      `json:"error_msg"`  // openapi 错误信息
}

// GetTemplateList
func GetTemplateList(params *GetTemplateListRequest) (*GetTemplateListResponsedata, error) {
	var (
		err        error
		defaultRet *GetTemplateListResponsedata
	)
	respData := &GetTemplateListResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/templatelist")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("offset", fmt.Sprintf("%v", params.Offset))
	client.AddGetParam("count", fmt.Sprintf("%v", params.Count))
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)

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
