package openapi

import (
	"fmt"
)

// GetTemplateLibraryListRequest 请求结构体
type GetTemplateLibraryListRequest struct {
	AccessToken string // 接口调用凭证
	Offset      int64  // 用于分页，表示从 offset 开始，默认值为 0
	Count       int64  // 用于分页，表示拉取 count 条记录，默认值为 0 ，最大为 20
}

// 响应结构体

type GetTemplateLibraryListResponsedatalistItem struct {
	ID    string `json:"id"`    // 模板标题 id（获取模板标题下的关键词库时需要）
	Title string `json:"title"` // 模板标题内容
}

type GetTemplateLibraryListResponsedata struct {
	List       []GetTemplateLibraryListResponsedatalistItem `json:"list"`        // 模板数组，对象中包含模板标题 id 、模板标题内容 title
	TotalCount int64                                        `json:"total_count"` // 模板库标题总数
}

type GetTemplateLibraryListResponse struct {
	Data      GetTemplateLibraryListResponsedata `json:"data"`       // 响应对象
	Errno     int64                              `json:"errno"`      // 错误码
	ErrMsg    string                             `json:"msg"`        // 错误信息
	ErrorCode int64                              `json:"error_code"` // openapi 错误码
	ErrorMsg  string                             `json:"error_msg"`  // openapi 错误信息
}

// GetTemplateLibraryList
func GetTemplateLibraryList(params *GetTemplateLibraryListRequest) (*GetTemplateLibraryListResponsedata, error) {
	var (
		err        error
		defaultRet *GetTemplateLibraryListResponsedata
	)
	respData := &GetTemplateLibraryListResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/librarylist")
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
