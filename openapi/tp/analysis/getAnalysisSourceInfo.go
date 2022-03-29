package analysis

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAnalysisSourceInfoRequest 请求结构体
type GetAnalysisSourceInfoRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	StartIndex  int64  // 数据偏移位置，接口默认返回 20 条数据，可使用该偏移量进行翻页查看
	StartDate   string // 开始日期：20190410
	EndDate     string // 结束日期：20190415
}

// 响应结构体

type GetAnalysisSourceInfoResponsedatadataListItem struct {
	AverageSessionTime  interface{} `json:"average_session_time"`   // 次均使用时长，平均每一次启动小程序的时间，等于总时长/总启动次数
	Name                interface{} `json:"name"`                   //
	NewUserCount        interface{} `json:"new_user_count"`         // 新用户数，90天内首次启动小程序的用户数
	NewUserScale        interface{} `json:"new_user_scale"`         // 新用户分布，某来源的新用户数占全来源的新用户总数的比例
	SessionCountPerUser interface{} `json:"session_count_per_user"` // 人均启动次数，小程序累计启动次数/启动用户数
	SessionScale        interface{} `json:"session_scale"`          // 启动次数分布，某来源的启动次数占同一层级来源的启动次数总和的比例
	UserCount           interface{} `json:"user_count"`             // 启动用户数，启动过小程序的用户（多次启动不重复计）
	UserScale           interface{} `json:"user_scale"`             // 启动用户分布，某来源的启动用户数占全来源的启动用户总数的比例
}

type GetAnalysisSourceInfoResponsedata struct {
	DataList []GetAnalysisSourceInfoResponsedatadataListItem `json:"data_list"` // 数据综合(指标内容)
	Offset   int64                                           `json:"offset"`    // 数据偏移位置
	TimeSpan string                                          `json:"time_span"` // 时间范围
	Total    int64                                           `json:"total"`     // 数据量
}

type GetAnalysisSourceInfoResponse struct {
	Data      GetAnalysisSourceInfoResponsedata `json:"data"`       // 响应参数
	Errno     int64                             `json:"errno"`      // 状态码
	ErrMsg    string                            `json:"msg"`        // 错误信息
	ErrorCode int64                             `json:"error_code"` // openapi 错误码
	ErrorMsg  string                            `json:"error_msg"`  // openapi 错误信息
}

// GetAnalysisSourceInfo
func GetAnalysisSourceInfo(params *GetAnalysisSourceInfoRequest) (*GetAnalysisSourceInfoResponsedata, error) {
	var (
		err        error
		defaultRet *GetAnalysisSourceInfoResponsedata
	)
	respData := &GetAnalysisSourceInfoResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/data/getanalysissource")
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
