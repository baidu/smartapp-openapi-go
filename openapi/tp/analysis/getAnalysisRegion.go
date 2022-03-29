package analysis

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAnalysisRegionRequest 请求结构体
type GetAnalysisRegionRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	StartIndex  int64  // 数据偏移位置，接口默认返回 20 条数据，可使用该偏移量进行翻页查看
	StartDate   string // 开始日期：20190410
	EndDate     string // 结束日期：20190415
}

// 响应结构体

type GetAnalysisRegionResponsedatasum struct {
	AverageSessionTimeDistrict interface{} `json:"average_session_time_district"` // 次均使用时长，该地域内的用户平均每一次使用应用程序（session）的时间
	NewUserCountDistrictRatio  interface{} `json:"new_user_count_district_ratio"` // 新用户分布，该地域的新用户数占全地域的新用户之和的比例
	SessionCountDistrict       interface{} `json:"session_count_district"`        // 启动次数，该地域用户的启动次数
	SessionCountDistrictRatio  interface{} `json:"session_count_district_ratio"`  // 启动次数比例分布，该地域用户的启动次数占全地域用户的启动次数之和的比例
	UserCountDistrictRatio     interface{} `json:"user_count_district_ratio"`     // 启动用户分布，该地域的启动用户数占全地域的启动用户（跨地域不去重）的比例
}

type GetAnalysisRegionResponsedatadataListItem struct {
	AverageSessionTimeDistrict interface{} `json:"average_session_time_district"` // 次均使用时长，该地域内的用户平均每一次使用应用程序（session）的时间
	Name                       interface{} `json:"name"`                          //
	NewUserCountDistrictRatio  interface{} `json:"new_user_count_district_ratio"` // 新用户分布，该地域的新用户数占全地域的新用户之和的比例
	SessionCountDistrict       interface{} `json:"session_count_district"`        // 启动次数，该地域用户的启动次数
	SessionCountDistrictRatio  interface{} `json:"session_count_district_ratio"`  // 启动次数比例分布，该地域用户的启动次数占全地域用户的启动次数之和的比例
	UserCountDistrictRatio     interface{} `json:"user_count_district_ratio"`     // 启动用户分布，该地域的启动用户数占全地域的启动用户（跨地域不去重）的比例
}

type GetAnalysisRegionResponsedata struct {
	DataList []GetAnalysisRegionResponsedatadataListItem `json:"data_list"` // 数据列表(指标内容)
	Offset   int64                                       `json:"offset"`    // 数据偏移位置
	Sum      GetAnalysisRegionResponsedatasum            `json:"sum"`       // 数据综合(指标内容)
	TimeSpan string                                      `json:"time_span"` // 时间范围
	Total    int64                                       `json:"total"`     // 数据量
}

type GetAnalysisRegionResponse struct {
	Data      GetAnalysisRegionResponsedata `json:"data"`       // 响应参数
	Errno     int64                         `json:"errno"`      // 状态码
	ErrMsg    string                        `json:"msg"`        // 错误信息
	ErrorCode int64                         `json:"error_code"` // openapi 错误码
	ErrorMsg  string                        `json:"error_msg"`  // openapi 错误信息
}

// GetAnalysisRegion
func GetAnalysisRegion(params *GetAnalysisRegionRequest) (*GetAnalysisRegionResponsedata, error) {
	var (
		err        error
		defaultRet *GetAnalysisRegionResponsedata
	)
	respData := &GetAnalysisRegionResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/data/getanalysisregion")
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
