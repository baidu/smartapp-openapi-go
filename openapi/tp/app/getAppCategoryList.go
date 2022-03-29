package app

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAppCategoryListRequest 请求结构体
type GetAppCategoryListRequest struct {
	AccessToken  string // 授权小程序的接口调用凭据
	CategoryType string // 1.个人类型类目 2.企业类型类目 为2时可以查出全部类目
}

// 响应结构体

type GetAppCategoryListResponsedataItemsubItems struct {
	CategoryName  string `json:"category_name"`  // 子类目名称
	CategoryQuali string `json:"category_quali"` // 类目 资质要求
	CategoryType  int64  `json:"category_type"`  // 子类目类型 规则同父类目类型
	ID            int64  `json:"id"`             // 子类目 Id
	NeedQuali     int64  `json:"need_quali"`     // 类目是否需资质 1：需要 0：不需要
}

type GetAppCategoryListResponsedataItem struct {
	CategoryName string                                                `json:"category_name"` // 父类目名称
	CategoryType int64                                                 `json:"category_type"` // 父类目类型 1.个人服务 2.企业服务 主体为企业类型所有类目均可设置，主体为个人类型时只能设置个人服务
	ID           int64                                                 `json:"id"`            // 类目 Id
	SubItems     map[string]GetAppCategoryListResponsedataItemsubItems `json:"sub_items"`     // 子类目 key为类目id value为类目详情
}

type GetAppCategoryListResponse struct {
	Data      []GetAppCategoryListResponsedataItem `json:"data"`       // 响应参数
	Errno     int64                                `json:"errno"`      // 状态码
	ErrMsg    string                               `json:"msg"`        // 错误信息
	ErrorCode int64                                `json:"error_code"` // openapi 错误码
	ErrorMsg  string                               `json:"error_msg"`  // openapi 错误信息
}

// GetAppCategoryList
func GetAppCategoryList(params *GetAppCategoryListRequest) ([]GetAppCategoryListResponsedataItem, error) {
	var (
		err        error
		defaultRet []GetAppCategoryListResponsedataItem
	)
	respData := &GetAppCategoryListResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/app/category/list")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("category_type", params.CategoryType)
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
