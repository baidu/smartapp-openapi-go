package search

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetTopQueryURLInfoRequest 请求结构体
type GetTopQueryURLInfoRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	TimeSpan    int64  // 时间区间。可取值1、7、30，分别代表 1 天、7 天、30 天
	PageNum     int64  // 分页页数。
	PageSize    int64  // 分页大小。最大值为100
}

// 响应结构体

type GetTopQueryURLInfoResponsedatadetailItem struct {
	ClickCount   int64  `json:"click_count"`   // 点击量。
	DisplayCount int64  `json:"display_count"` // 展现量。
	Path         string `json:"path"`          // path路径。
	Rate         string `json:"rate"`          // 点展比(%), 保留两位小数。
	WebURL       string `json:"web_url"`       // web化url地址。
}

type GetTopQueryURLInfoResponsedata struct {
	Detail            []GetTopQueryURLInfoResponsedatadetailItem `json:"detail"`              // 日期纬度的点击与展示列表详情。
	Rate              string                                     `json:"rate"`                // 点展比(%), 保留两位小数。
	TotalClickCount   int64                                      `json:"total_click_count"`   // 累计点击量。
	TotalDisplayCount int64                                      `json:"total_display_count"` // 累计展现量。
	TotalNum          int64                                      `json:"total_num"`           //
}

type GetTopQueryURLInfoResponse struct {
	Data      GetTopQueryURLInfoResponsedata `json:"data"`       // 响应参数
	Errno     int64                          `json:"errno"`      // 状态码
	ErrMsg    string                         `json:"msg"`        // 错误信息
	ErrorCode int64                          `json:"error_code"` // openapi 错误码
	ErrorMsg  string                         `json:"error_msg"`  // openapi 错误信息
}

// GetTopQueryURLInfo
func GetTopQueryURLInfo(params *GetTopQueryURLInfoRequest) (*GetTopQueryURLInfoResponsedata, error) {
	var (
		err        error
		defaultRet *GetTopQueryURLInfoResponsedata
	)
	respData := &GetTopQueryURLInfoResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/promotion/topquery/info")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("time_span", fmt.Sprintf("%v", params.TimeSpan))
	client.AddGetParam("page_num", fmt.Sprintf("%v", params.PageNum))
	client.AddGetParam("page_size", fmt.Sprintf("%v", params.PageSize))
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
