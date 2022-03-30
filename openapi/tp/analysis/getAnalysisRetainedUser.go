package analysis

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAnalysisRetainedUserRequest 请求结构体
type GetAnalysisRetainedUserRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	ReportType  string // 报告数据类型:可选count数量类型、ratio比率类型，决定返回的指标内容
	StartIndex  int64  // 数据偏移位置，接口默认返回 20 条数据，可使用该偏移量进行翻页查看
	StartDate   string // 开始日期：20190410
	EndDate     string // 结束日期：20190415
	Gran        string // 数据粒度：day/week/month
}

// 响应结构体

type GetAnalysisRetainedUserResponsedatasum struct {
	ActiveUserCount   interface{} `json:"active_user_count"`   // 活跃用户数量，某日（周、月）启动过小程序的用户数（多次启动不重复计）
	RetainedRatio     interface{} `json:"retained_ratio"`      // 留存用户率，留存用户数占新用户数的比例
	RetainedUserCount interface{} `json:"retained_user_count"` // 留存用户数量，某日（周、月）新用户（或启动用户）在目标时间段再次启动小程序的用户数
}

type GetAnalysisRetainedUserResponsedata struct {
	DataList []map[string]interface{}               `json:"data_list"` // 数据列表(指标内容)
	Offset   int64                                  `json:"offset"`    // 数据偏移位置
	Sum      GetAnalysisRetainedUserResponsedatasum `json:"sum"`       // 数据综合(指标内容)
	TimeSpan string                                 `json:"time_span"` // 时间范围
	Total    int64                                  `json:"total"`     // 数据量
}

type GetAnalysisRetainedUserResponse struct {
	Data      GetAnalysisRetainedUserResponsedata `json:"data"`       // 响应参数
	Errno     int64                               `json:"errno"`      // 状态码
	ErrMsg    string                              `json:"msg"`        // 错误信息
	ErrorCode int64                               `json:"error_code"` // openapi 错误码
	ErrorMsg  string                              `json:"error_msg"`  // openapi 错误信息
}

// GetAnalysisRetainedUser
func GetAnalysisRetainedUser(params *GetAnalysisRetainedUserRequest) (*GetAnalysisRetainedUserResponsedata, error) {
	var (
		err        error
		defaultRet *GetAnalysisRetainedUserResponsedata
	)
	respData := &GetAnalysisRetainedUserResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/data/getanalysisretaineduser")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("report_type", params.ReportType)
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
