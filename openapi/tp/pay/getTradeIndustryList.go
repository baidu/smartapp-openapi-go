package pay

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetTradeIndustryListRequest 请求结构体
type GetTradeIndustryListRequest struct {
	AccessToken string // 第三方平台的接口调用凭据
}

// 响应结构体

type GetTradeIndustryListResponsedataItemindustryListItemindustryListItem struct {
	IndustryID       int64  `json:"industryId"`       //
	IndustryName     string `json:"industryName"`     //
	NeedPermit       int64  `json:"needPermit"`       //
	ParentIndustryID int64  `json:"parentIndustryId"` //
	PermitDesc       string `json:"permitDesc"`       //
}

type GetTradeIndustryListResponsedataItemindustryListItem struct {
	IndustryID       int64                                                                  `json:"industryId"`       //
	IndustryList     []GetTradeIndustryListResponsedataItemindustryListItemindustryListItem `json:"industryList"`     //
	IndustryName     string                                                                 `json:"industryName"`     //
	NeedPermit       int64                                                                  `json:"needPermit"`       //
	ParentIndustryID int64                                                                  `json:"parentIndustryId"` //
	PermitDesc       string                                                                 `json:"permitDesc"`       //
}

type GetTradeIndustryListResponsedataItem struct {
	IndustryID       int64                                                  `json:"industryId"`       //
	IndustryList     []GetTradeIndustryListResponsedataItemindustryListItem `json:"industryList"`     //
	IndustryName     string                                                 `json:"industryName"`     //
	NeedPermit       int64                                                  `json:"needPermit"`       //
	ParentIndustryID int64                                                  `json:"parentIndustryId"` //
	PermitDesc       string                                                 `json:"permitDesc"`       //
}

type GetTradeIndustryListResponse struct {
	Data      []GetTradeIndustryListResponsedataItem `json:"data"`       // 响应参数
	Errno     int64                                  `json:"errno"`      // 状态码
	ErrMsg    string                                 `json:"msg"`        // 错误信息
	ErrorCode int64                                  `json:"error_code"` // openapi 错误码
	ErrorMsg  string                                 `json:"error_msg"`  // openapi 错误信息
}

// GetTradeIndustryList
func GetTradeIndustryList(params *GetTradeIndustryListRequest) ([]GetTradeIndustryListResponsedataItem, error) {
	var (
		err        error
		defaultRet []GetTradeIndustryListResponsedataItem
	)
	respData := &GetTradeIndustryListResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/tp/getTradeIndustryList")
	client.AddGetParam("access_token", params.AccessToken)
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
	return respData.Data, nil
}
