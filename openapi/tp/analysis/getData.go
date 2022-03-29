package analysis

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetDataRequest 请求结构体
type GetDataRequest struct {
	AccessToken string      // 第三方平台的接口调用凭据
	Scene       interface{} // 小程序来源ID (场景值) 。不传则查询所有场景,场景值参数参考:百度APP 场景值
	Metrics     string      // 指标以逗号分隔
	StartDate   string      // 起始时间戳,格式如 20190321
	EndDate     string      // 结束时间戳,格式如 20190325
	StartIndex  interface{} // 偏移量,默认为0(分页操作从第几条开始展示)
	MaxResults  interface{} // 页面大小,默认值20(分页操作查询条数)
}

// 响应结构体

type GetDataResponsedataresultItem struct {
	Date                   string `json:"date"`                       // 日期
	TpActivityDegree       string `json:"tp_activity_degree	"`        // 小程序活跃度
	TpDayAppCount          string `json:"tp_day_app_count	"`          // 当日小程序数量
	TpDayBackFlowCount     string `json:"tp_day_back_flow_count"`     // 回流次数
	TpDayNewUserCount      string `json:"tp_day_new_user_count"`      // 当日启动新用户数
	TpDayPageCount         string `json:"tp_day_page_count"`          // 当日页面访问次数
	TpDaySessionCount      string `json:"tp_day_session_count"`       // 当日启动次数
	TpDayShareCount        string `json:"tp_day_share_count"`         // 分享次数
	TpDayUserCount         string `json:"tp_day_user_count	"`         // 当日启动用户数
	TpMonthAppCount        string `json:"tp_month_app_count"`         // 近30日小程序数量
	TpMonthNewUserCount    string `json:"tp_month_new_user_count"`    // 近30日启动新用户数
	TpMonthSessionCount    string `json:"tp_month_session_count"`     // 近30日启动次数
	TpMonthUserCount       string `json:"tp_month_user_count	"`       // 近30日启动用户数
	TpSessionTimePerDay    string `json:"tp_session_time_per_day"`    // 当日次均使用时长
	TpSessionTimePerPerson string `json:"tp_session_time_per_person"` // 当日人均使用时长
	TpWeekAppCount         string `json:"tp_week_app_count"`          // 近7日小程序数量
	TpWeekNewUserCount     string `json:"tp_week_new_user_count"`     // 近7日启动新用户数
	TpWeekSessionCount     string `json:"tp_week_session_count	"`     // 近7日启动次数
	TpWeekUserCount        string `json:"tp_week_user_count	"`        // 近7日启动用户数
}

type GetDataResponsedata struct {
	Offset int64                           `json:"offset"` // 数据偏移位置
	Result []GetDataResponsedataresultItem `json:"result"` //
	Total  int64                           `json:"total"`  // 数据量
}

type GetDataResponse struct {
	Data      GetDataResponsedata `json:"data"`       // 响应参数
	Errno     int64               `json:"errno"`      // 状态码
	ErrMsg    string              `json:"msg"`        // 错误信息
	ErrorCode int64               `json:"error_code"` // openapi 错误码
	ErrorMsg  string              `json:"error_msg"`  // openapi 错误信息
}

// GetData
func GetData(params *GetDataRequest) (*GetDataResponsedata, error) {
	var (
		err        error
		defaultRet *GetDataResponsedata
	)
	respData := &GetDataResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/data/gettpdata")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("scene", params.Scene)
	client.AddGetParam("metrics", params.Metrics)
	client.AddGetParam("start_date", params.StartDate)
	client.AddGetParam("end_date", params.EndDate)
	client.AddGetParam("start_index", params.StartIndex)
	client.AddGetParam("max_results", params.MaxResults)
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
