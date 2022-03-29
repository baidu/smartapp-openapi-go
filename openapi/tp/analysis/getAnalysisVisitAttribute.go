package analysis

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAnalysisVisitAttributeRequest 请求结构体
type GetAnalysisVisitAttributeRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	StartDate   string // 开始日期：20190410
	EndDate     string // 结束日期：20190415
}

// 响应结构体

type GetAnalysisVisitAttributeResponsedatadataListItemageItem struct {
	AttributeID interface{} `json:"attribute_id"` //
	CountRatio  interface{} `json:"count_ratio"`  //
	Name        interface{} `json:"name"`         //
}

type GetAnalysisVisitAttributeResponsedatadataListItemsexItem struct {
	AttributeID interface{} `json:"attribute_id"` //
	CountRatio  interface{} `json:"count_ratio"`  //
	Name        interface{} `json:"name"`         //
}

type GetAnalysisVisitAttributeResponsedatadataListIteminterestItem struct {
	AttributeID interface{} `json:"attribute_id"` //
	CountRatio  interface{} `json:"count_ratio"`  //
	Name        interface{} `json:"name"`         //
}

type GetAnalysisVisitAttributeResponsedatadataListItem struct {
	Age      []GetAnalysisVisitAttributeResponsedatadataListItemageItem      `json:"age"`      // 年龄分布
	Interest []GetAnalysisVisitAttributeResponsedatadataListIteminterestItem `json:"interest"` // 兴趣分布
	Sex      []GetAnalysisVisitAttributeResponsedatadataListItemsexItem      `json:"sex"`      // 年龄分布
}

type GetAnalysisVisitAttributeResponsedata struct {
	DataList []GetAnalysisVisitAttributeResponsedatadataListItem `json:"data_list"` //
}

type GetAnalysisVisitAttributeResponse struct {
	Data      GetAnalysisVisitAttributeResponsedata `json:"data"`       // 响应参数
	Errno     int64                                 `json:"errno"`      // 状态码
	ErrMsg    string                                `json:"msg"`        // 错误信息
	ErrorCode int64                                 `json:"error_code"` // openapi 错误码
	ErrorMsg  string                                `json:"error_msg"`  // openapi 错误信息
}

// GetAnalysisVisitAttribute
func GetAnalysisVisitAttribute(params *GetAnalysisVisitAttributeRequest) (*GetAnalysisVisitAttributeResponsedata, error) {
	var (
		err        error
		defaultRet *GetAnalysisVisitAttributeResponsedata
	)
	respData := &GetAnalysisVisitAttributeResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/data/getanalysisvisitattribute")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("start_date", params.StartDate)
	client.AddPostParam("end_date", params.EndDate)

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
