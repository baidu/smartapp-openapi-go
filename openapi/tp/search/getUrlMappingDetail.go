package search

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetURLMappingDetailRequest 请求结构体
type GetURLMappingDetailRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	RuleID      int64  // 规则ID
}

// 响应结构体

type GetURLMappingDetailResponsedataItem struct {
	AdapterNum    int64  `json:"adapterNum"`    // 适配topurl的数量
	AdapterStatus int64  `json:"adapterStatus"` // 规则状态
	AppPattern    string `json:"appPattern"`    // 小程序Path表达式
	CoverNum      int64  `json:"coverNum"`      // 累计替换 URL 量，即累计截止 t-1已完成替换的的 url 数量
	CreateMethod  int64  `json:"createMethod"`  //
	H5Pattern     string `json:"h5Pattern"`     // 移动url表达式
	IsFetch       bool   `json:"isFetch"`       // 是否匹配topurl生成path
	MatchNum      int64  `json:"matchNum"`      // 匹配topurl的数量
	Msg           string `json:"msg"`           //
	RuleID        int64  `json:"ruleId"`        // 规则id
	Status        int64  `json:"status"`        //
	SubmitTime    string `json:"submitTime"`    //
	URLRate       int64  `json:"urlRate"`       // top 流量 URL 替换占比=累计截止 t-1已完成替换的top url 数量/总 top url 数量
}

type GetURLMappingDetailResponse struct {
	Data      []GetURLMappingDetailResponsedataItem `json:"data"`       // 响应参数
	Errno     int64                                 `json:"errno"`      // 状态码
	ErrMsg    string                                `json:"msg"`        // 错误信息
	ErrorCode int64                                 `json:"error_code"` // openapi 错误码
	ErrorMsg  string                                `json:"error_msg"`  // openapi 错误信息
}

// GetURLMappingDetail
func GetURLMappingDetail(params *GetURLMappingDetailRequest) ([]GetURLMappingDetailResponsedataItem, error) {
	var (
		err        error
		defaultRet []GetURLMappingDetailResponsedataItem
	)
	respData := &GetURLMappingDetailResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/flow/geturlmappingdetail")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("rule_id", fmt.Sprintf("%v", params.RuleID))

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
	return respData.Data, nil
}
