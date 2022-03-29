package analysis

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAnalysisUserTrendRequest 请求结构体
type GetAnalysisUserTrendRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	StartIndex  int64  // 数据偏移位置，接口默认返回 20 条数据，可使用该偏移量进行翻页查看
	StartDate   string // 开始日期：20190410
	EndDate     string // 结束日期：20190415
	Gran        string // 数据粒度：day/hour/week/month
}

// 响应结构体

type GetAnalysisUserTrendResponsedatasum struct {
	AverageSessionTime   interface{} `json:"average_session_time"`    // 次均使用时长，平均每一次启动小程序的时间，等于总时长/总启动次数
	NewUserCount         interface{} `json:"new_user_count"`          // 新用户数，当日首次启动小程序的用户数
	NewUserRatio         interface{} `json:"new_user_ratio"`          // 新用户占比，新用户数占启动用户总数的比例
	OldUserCount         interface{} `json:"old_user_count"`          // 老用户数，当日启动用户中，以前也启动过小程序的用户
	OldUserRatio         interface{} `json:"old_user_ratio"`          // 老用户占比，当日老用户占总的启动用户的比例
	SessionCount         interface{} `json:"session_count"`           // 启动次数，启动小程序的总次数。"一次启动"是指用户打开小程序到主动退出（或超时退出）为止。
	SessionTimePerPerson interface{} `json:"session_time_per_person"` // 人均使用时长，平均每个用户使用应用程序的时间，等于总时长/总启动用户数
	UserCount            interface{} `json:"user_count"`              // 启动用户，启动过小程序的用户数（多次启动不重复计）
}

type GetAnalysisUserTrendResponsedatadataListItem struct {
	AverageSessionTime   interface{} `json:"average_session_time"`    // 次均使用时长，平均每一次启动小程序的时间，等于总时长/总启动次数
	Name                 interface{} `json:"name"`                    //
	NewUserCount         interface{} `json:"new_user_count"`          // 新用户数，当日首次启动小程序的用户数
	NewUserRatio         interface{} `json:"new_user_ratio"`          // 新用户占比，新用户数占启动用户总数的比例
	OldUserCount         interface{} `json:"old_user_count"`          // 老用户数，当日启动用户中，以前也启动过小程序的用户
	OldUserRatio         interface{} `json:"old_user_ratio"`          // 老用户占比，当日老用户占总的启动用户的比例
	SessionCount         interface{} `json:"session_count"`           // 启动次数，启动小程序的总次数。"一次启动"是指用户打开小程序到主动退出（或超时退出）为止。
	SessionTimePerPerson interface{} `json:"session_time_per_person"` // 人均使用时长，平均每个用户使用应用程序的时间，等于总时长/总启动用户数
	UserCount            interface{} `json:"user_count"`              // 启动用户，启动过小程序的用户数（多次启动不重复计）
}

type GetAnalysisUserTrendResponsedata struct {
	DataList []GetAnalysisUserTrendResponsedatadataListItem `json:"data_list"` // 数据列表(指标内容)
	Offset   int64                                          `json:"offset"`    // 数据偏移位置
	Sum      GetAnalysisUserTrendResponsedatasum            `json:"sum"`       // 数据综合(指标内容)
	TimeSpan string                                         `json:"time_span"` // 时间范围
	Total    int64                                          `json:"total"`     // 数据量
}

type GetAnalysisUserTrendResponse struct {
	Data      GetAnalysisUserTrendResponsedata `json:"data"`       // 响应参数
	Errno     int64                            `json:"errno"`      // 状态码
	ErrMsg    string                           `json:"msg"`        // 错误信息
	ErrorCode int64                            `json:"error_code"` // openapi 错误码
	ErrorMsg  string                           `json:"error_msg"`  // openapi 错误信息
}

// GetAnalysisUserTrend
func GetAnalysisUserTrend(params *GetAnalysisUserTrendRequest) (*GetAnalysisUserTrendResponsedata, error) {
	var (
		err        error
		defaultRet *GetAnalysisUserTrendResponsedata
	)
	respData := &GetAnalysisUserTrendResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/data/getanalysisusertrend")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("start_index", fmt.Sprintf("%v", params.StartIndex))
	client.AddPostParam("start_date", params.StartDate)
	client.AddPostParam("end_date", params.EndDate)
	client.AddPostParam("gran", params.Gran)

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
