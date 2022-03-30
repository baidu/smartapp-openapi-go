package search

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetSitemapMetaRequest 请求结构体
type GetSitemapMetaRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	TimeSpan    int64  // 时间区间。可取值1、7、30，分别代表 1 天、7 天、30 天
}

// 响应结构体

type GetSitemapMetaResponsedatatrafficDataItem struct {
	ClickCount     int64  `json:"click_count"`      // 点击量。
	ClickViewRatio string `json:"click_view_ratio"` // 点展比(%), 保留两位小数。
	Date           string `json:"date"`             // 日期。
	DisplayCount   int64  `json:"display_count"`    // 展现量。
}

type GetSitemapMetaResponsedata struct {
	ClickDisplayRatio string                                      `json:"click_display_ratio"` // 点展比。
	TotalClickCount   int64                                       `json:"total_click_count"`   // 累计点击量。
	TotalDisplayCount int64                                       `json:"total_display_count"` // 累计展现量。
	TrafficData       []GetSitemapMetaResponsedatatrafficDataItem `json:"traffic_data"`        // 日期纬度的点击与展示列表详情。
}

type GetSitemapMetaResponse struct {
	Data      GetSitemapMetaResponsedata `json:"data"`       // 响应参数
	Errno     int64                      `json:"errno"`      // 状态码
	ErrMsg    string                     `json:"msg"`        // 错误信息
	ErrorCode int64                      `json:"error_code"` // openapi 错误码
	ErrorMsg  string                     `json:"error_msg"`  // openapi 错误信息
}

// GetSitemapMeta
func GetSitemapMeta(params *GetSitemapMetaRequest) (*GetSitemapMetaResponsedata, error) {
	var (
		err        error
		defaultRet *GetSitemapMetaResponsedata
	)
	respData := &GetSitemapMetaResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/promotion/sitemap/getanalysisctr")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("time_span", fmt.Sprintf("%v", params.TimeSpan))
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
