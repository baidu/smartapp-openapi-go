package analysis

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAnalysisVisitPageRequest 请求结构体
type GetAnalysisVisitPageRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	StartIndex  int64  // 数据偏移位置，接口默认返回 20 条数据，可使用该偏移量进行翻页查看
	StartDate   string // 开始日期：20190410
	EndDate     string // 结束日期：20190415
}

// 响应结构体

type GetAnalysisVisitPageResponsedatasum struct {
	AverageStayTime interface{} `json:"average_stay_time"` // 次均访问时长，用户访问当前页面的次均访问时间
	BounceRatio     interface{} `json:"bounce_ratio"`      // 退出率，用户从当前页面离开小程序的比例
	EntryCount      interface{} `json:"entry_count"`       // 入口页次数，该页面作为启动小程序时第一个访问的页面的启动次数
	ExitCount       interface{} `json:"exit_count"`        // 退出页次数，该页面作为关闭小程序时最后一个访问的页面的启动次数
	PvCount         interface{} `json:"pv_count"`          // 页面访问次数，页面被访问的次数，多次跳转重复访问也会被计入
	PvRatio         interface{} `json:"pv_ratio"`          // 访问次数占比，当前页面访问次数占全部页面访问次数的比例
	PvUserCount     interface{} `json:"pv_user_count"`     // 访问用户数，访问当前页面的总用户数
	StayTimeRatio   interface{} `json:"stay_time_ratio"`   // 访问时长占比，用户访问当前页面的访问时长的总和占用户在全部页面的访问时长总和的比例
}

type GetAnalysisVisitPageResponsedatadataListItem struct {
	AverageStayTime interface{} `json:"average_stay_time"` // 次均访问时长，用户访问当前页面的次均访问时间
	BounceRatio     interface{} `json:"bounce_ratio"`      // 退出率，用户从当前页面离开小程序的比例
	EntryCount      interface{} `json:"entry_count"`       // 入口页次数，该页面作为启动小程序时第一个访问的页面的启动次数
	ExitCount       interface{} `json:"exit_count"`        // 退出页次数，该页面作为关闭小程序时最后一个访问的页面的启动次数
	PageAlias       interface{} `json:"pageAlias"`         //
	PageID          interface{} `json:"pageId"`            //
	PageName        interface{} `json:"pageName"`          //
	PvCount         interface{} `json:"pv_count"`          // 页面访问次数，页面被访问的次数，多次跳转重复访问也会被计入
	PvRatio         interface{} `json:"pv_ratio"`          // 访问次数占比，当前页面访问次数占全部页面访问次数的比例
	PvUserCount     interface{} `json:"pv_user_count"`     // 访问用户数，访问当前页面的总用户数
	StayTimeRatio   interface{} `json:"stay_time_ratio"`   // 访问时长占比，用户访问当前页面的访问时长的总和占用户在全部页面的访问时长总和的比例
}

type GetAnalysisVisitPageResponsedata struct {
	DataList []GetAnalysisVisitPageResponsedatadataListItem `json:"data_list"` // 数据列表(指标内容)
	Offset   int64                                          `json:"offset"`    // 数据偏移位置
	Sum      GetAnalysisVisitPageResponsedatasum            `json:"sum"`       // 数据综合(指标内容)
	TimeSpan string                                         `json:"time_span"` // 时间范围
	Total    int64                                          `json:"total"`     // 数据量
}

type GetAnalysisVisitPageResponse struct {
	Data      GetAnalysisVisitPageResponsedata `json:"data"`       // 响应参数
	Errno     int64                            `json:"errno"`      // 状态码
	ErrMsg    string                           `json:"msg"`        // 错误信息
	ErrorCode int64                            `json:"error_code"` // openapi 错误码
	ErrorMsg  string                           `json:"error_msg"`  // openapi 错误信息
}

// GetAnalysisVisitPage
func GetAnalysisVisitPage(params *GetAnalysisVisitPageRequest) (*GetAnalysisVisitPageResponsedata, error) {
	var (
		err        error
		defaultRet *GetAnalysisVisitPageResponsedata
	)
	respData := &GetAnalysisVisitPageResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/data/getanalysisvisitpage")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("start_index", fmt.Sprintf("%v", params.StartIndex))
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
