package search

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// QueryURLMappingRequest 请求结构体
type QueryURLMappingRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	PageNum     int64  // 分页页码。从1开始
	PageSize    int64  // 分页大小。最大值为 100
}

// 响应结构体

type QueryURLMappingResponsedataruleDetailListItem struct {
	AppRule      string `json:"appRule"`      // 小程序rule
	CreateMethod int64  `json:"createMethod"` // 创建方式
	H5Rule       string `json:"h5Rule"`       // h5规则
	Msg          string `json:"msg"`          // 规则状态说明
	RuleID       int64  `json:"ruleId"`       // 规则ID
	Status       int64  `json:"status"`       // 规则状态
	SubmitTime   string `json:"submitTime"`   // 规则提交时间
}

type QueryURLMappingResponsedata struct {
	RuleDetailList   []QueryURLMappingResponsedataruleDetailListItem `json:"ruleDetailList"`   // 规则列表
	RuleRemainingNum int64                                           `json:"ruleRemainingNum"` // 规则剩余可用数量
	RuleTotalNum     int64                                           `json:"ruleTotalNum"`     // 规则总数
	TotalAdapterNum  int64                                           `json:"totalAdapterNum"`  // 适配资源生效量
}

type QueryURLMappingResponse struct {
	Data      QueryURLMappingResponsedata `json:"data"`       // 响应参数
	Errno     int64                       `json:"errno"`      // 状态码
	ErrMsg    string                      `json:"msg"`        // 错误信息
	ErrorCode int64                       `json:"error_code"` // openapi 错误码
	ErrorMsg  string                      `json:"error_msg"`  // openapi 错误信息
}

// QueryURLMapping
func QueryURLMapping(params *QueryURLMappingRequest) (*QueryURLMappingResponsedata, error) {
	var (
		err        error
		defaultRet *QueryURLMappingResponsedata
	)
	respData := &QueryURLMappingResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/flow/queryurlmapping")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("page_num", fmt.Sprintf("%v", params.PageNum))
	client.AddPostParam("page_size", fmt.Sprintf("%v", params.PageSize))

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
