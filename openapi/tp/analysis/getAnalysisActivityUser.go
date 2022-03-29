package analysis

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAnalysisActivityUserRequest 请求结构体
type GetAnalysisActivityUserRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	StartIndex  int64  // 数据偏移位置，接口默认返回 20 条数据，可使用该偏移量进行翻页查看
	StartDate   string // 开始日期：20190410
	EndDate     string // 结束日期：20190415
}

// 响应结构体

type GetAnalysisActivityUserResponsedatasum struct {
	AccumulativeUserCount  interface{} `json:"accumulative_user_count"`   // 累计启动用户，累计启动过小程序的用户数量（多次启动不重复计）
	DailyActivityDegree    interface{} `json:"daily_activity_degree"`     // 日活跃度，当日启动用户/ 累计用户
	DayMonthActivityDegree interface{} `json:"day_month_activity_degree"` // 日活/月活，日启动用户数/月启动用户数，反映当日用户活跃与近30日的比较水平
	LostRatio              interface{} `json:"lost_ratio"`                // 流失率，流失用户/累计用户
	LostUserCount          interface{} `json:"lost_user_count"`           // 流失用户，最近60天（含查询当日）没有启动过小程序的用户（已去重）
	MonthlyActivityDegree  interface{} `json:"monthly_activity_degree"`   // 月活跃度，月活跃用户/累计用户
	MonthlyUserCount       interface{} `json:"monthly_user_count"`        // 月活跃用户，最近30天（含查询当日）启动过小程序的用户数（多次启动不重复计）
	UserCount              interface{} `json:"user_count"`                // 启动用户，启动过小程序的用户数（多次启动不重复计）
	WeeklyActivityDegree   interface{} `json:"weekly_activity_degree"`    // 周活跃度，周活跃用户/累计用户
	WeeklyUserCount        interface{} `json:"weekly_user_count"`         // 周活跃用户，最近7天（含查询当日）启动过小程序的用户数（多次启动不重复计）
}

type GetAnalysisActivityUserResponsedatadataListItem struct {
	AccumulativeUserCount  interface{} `json:"accumulative_user_count"`   // 累计启动用户，累计启动过小程序的用户数量（多次启动不重复计）
	DailyActivityDegree    interface{} `json:"daily_activity_degree"`     // 日活跃度，当日启动用户/ 累计用户
	DayMonthActivityDegree interface{} `json:"day_month_activity_degree"` // 日活/月活，日启动用户数/月启动用户数，反映当日用户活跃与近30日的比较水平
	LostRatio              interface{} `json:"lost_ratio"`                // 流失率，流失用户/累计用户
	LostUserCount          interface{} `json:"lost_user_count"`           // 流失用户，最近60天（含查询当日）没有启动过小程序的用户（已去重）
	MonthlyActivityDegree  interface{} `json:"monthly_activity_degree"`   // 月活跃度，月活跃用户/累计用户
	MonthlyUserCount       interface{} `json:"monthly_user_count"`        // 月活跃用户，最近30天（含查询当日）启动过小程序的用户数（多次启动不重复计）
	Name                   interface{} `json:"name"`                      //
	UserCount              interface{} `json:"user_count"`                // 启动用户，启动过小程序的用户数（多次启动不重复计）
	WeeklyActivityDegree   interface{} `json:"weekly_activity_degree"`    // 周活跃度，周活跃用户/累计用户
	WeeklyUserCount        interface{} `json:"weekly_user_count"`         // 周活跃用户，最近7天（含查询当日）启动过小程序的用户数（多次启动不重复计）
}

type GetAnalysisActivityUserResponsedata struct {
	DataList []GetAnalysisActivityUserResponsedatadataListItem `json:"data_list"` // 数据列表(指标内容)
	Offset   int64                                             `json:"offset"`    // 数据偏移位置
	Sum      GetAnalysisActivityUserResponsedatasum            `json:"sum"`       // 数据综合(指标内容)
	TimeSpan string                                            `json:"time_span"` // 时间范围
	Total    int64                                             `json:"total"`     // 数据量
}

type GetAnalysisActivityUserResponse struct {
	Data      GetAnalysisActivityUserResponsedata `json:"data"`       // 响应参数
	Errno     int64                               `json:"errno"`      // 状态码
	ErrMsg    string                              `json:"msg"`        // 错误信息
	ErrorCode int64                               `json:"error_code"` // openapi 错误码
	ErrorMsg  string                              `json:"error_msg"`  // openapi 错误信息
}

// GetAnalysisActivityUser
func GetAnalysisActivityUser(params *GetAnalysisActivityUserRequest) (*GetAnalysisActivityUserResponsedata, error) {
	var (
		err        error
		defaultRet *GetAnalysisActivityUserResponsedata
	)
	respData := &GetAnalysisActivityUserResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/data/getanalysisactivityuser")
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
