package msgtemplate

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetMsgTemplateLibraryListRequest 请求结构体
type GetMsgTemplateLibraryListRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	Offset      int64  // 用于分页，表示从offset开始，默认值为0
	Count       int64  // 用于分页，表示拉取count条记录，默认值为0,最大为20
}

// 响应结构体

type GetMsgTemplateLibraryListResponsedatalistItem struct {
	ID    string `json:"id"`    // 模板标题id（获取模板标题下的关键词库时需要）
	Title string `json:"title"` // 模板标题内容
}

type GetMsgTemplateLibraryListResponsedata struct {
	List       []GetMsgTemplateLibraryListResponsedatalistItem `json:"list"`        //
	TotalCount int64                                           `json:"total_count"` // 模板库标题总数
}

type GetMsgTemplateLibraryListResponse struct {
	Data      GetMsgTemplateLibraryListResponsedata `json:"data"`       // 响应参数
	Errno     int64                                 `json:"errno"`      // 状态码
	ErrMsg    string                                `json:"msg"`        // 错误信息
	ErrorCode int64                                 `json:"error_code"` // openapi 错误码
	ErrorMsg  string                                `json:"error_msg"`  // openapi 错误信息
}

// GetMsgTemplateLibraryList
func GetMsgTemplateLibraryList(params *GetMsgTemplateLibraryListRequest) (*GetMsgTemplateLibraryListResponsedata, error) {
	var (
		err        error
		defaultRet *GetMsgTemplateLibraryListResponsedata
	)
	respData := &GetMsgTemplateLibraryListResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/library/list")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("offset", fmt.Sprintf("%v", params.Offset))
	client.AddGetParam("count", fmt.Sprintf("%v", params.Count))
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)

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
