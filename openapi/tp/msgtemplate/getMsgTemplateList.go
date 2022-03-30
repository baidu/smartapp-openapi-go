package msgtemplate

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetMsgTemplateListRequest 请求结构体
type GetMsgTemplateListRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	Offset      int64  // 用于分页，表示从offset开始，默认值为0
	Count       int64  // 用于分页，表示拉取count条记录，默认值为0,最大为20
}

// 响应结构体

type GetMsgTemplateListResponsedatalistItem struct {
	Content    string `json:"content"`     // 模板内容
	Example    string `json:"example"`     // 模板内容示例
	TemplateID string `json:"template_id"` // 模板 id ，发送小程序模板消息时所需
	Title      string `json:"title"`       // 模板标题
}

type GetMsgTemplateListResponsedata struct {
	List       []GetMsgTemplateListResponsedatalistItem `json:"list"`        // 模板对象数组，对象中包含模板 id 、模板标题 title 、模板内容 content 、模板内容示例 example
	TotalCount int64                                    `json:"total_count"` // 模板库标题总数
}

type GetMsgTemplateListResponse struct {
	Data      GetMsgTemplateListResponsedata `json:"data"`       // 响应参数
	Errno     int64                          `json:"errno"`      // 状态码
	ErrMsg    string                         `json:"msg"`        // 错误信息
	ErrorCode int64                          `json:"error_code"` // openapi 错误码
	ErrorMsg  string                         `json:"error_msg"`  // openapi 错误信息
}

// GetMsgTemplateList
func GetMsgTemplateList(params *GetMsgTemplateListRequest) (*GetMsgTemplateListResponsedata, error) {
	var (
		err        error
		defaultRet *GetMsgTemplateListResponsedata
	)
	respData := &GetMsgTemplateListResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/list")
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
