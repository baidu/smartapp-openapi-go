package grade

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// AppGradeRequest 请求结构体
type AppGradeRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	PageNo      int64  // 分页页码（从 1 开始）
}

// 响应结构体

type AppGradeResponsedatadataListItem struct {
	AppID  int64  `json:"app_id"`  // 小程序 id
	AppKey string `json:"app_key"` // 小程序唯一识别码
	Grade  string `json:"grade"`   // 小程序等级 · 未评级 · A 级 · B 级 · C 级
}

type AppGradeResponsedata struct {
	DataList  []AppGradeResponsedatadataListItem `json:"dataList"`  //
	PageNo    int64                              `json:"pageNo"`    // 页码
	PageSize  int64                              `json:"pageSize"`  // 每页展示数据
	Total     int64                              `json:"total"`     // 总数据量
	TotalPage int64                              `json:"totalPage"` // 总页数
}

type AppGradeResponse struct {
	Data      AppGradeResponsedata `json:"data"`       // 响应参数
	Errno     int64                `json:"errno"`      // 状态码
	ErrMsg    string               `json:"msg"`        // 错误信息
	ErrorCode int64                `json:"error_code"` // openapi 错误码
	ErrorMsg  string               `json:"error_msg"`  // openapi 错误信息
}

// AppGrade
func AppGrade(params *AppGradeRequest) (*AppGradeResponsedata, error) {
	var (
		err        error
		defaultRet *AppGradeResponsedata
	)
	respData := &AppGradeResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/app/grade")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("page_no", fmt.Sprintf("%v", params.PageNo))
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
