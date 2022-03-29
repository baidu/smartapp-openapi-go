package analysis

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAnalysisTerminalInfoRequest 请求结构体
type GetAnalysisTerminalInfoRequest struct {
	AccessToken  string // 授权小程序的接口调用凭据
	TerminalType string // 终端数据类型
	StartIndex   int64  // 数据偏移位置，接口默认返回 20 条数据，可使用该偏移量进行翻页查看
	StartDate    string // 开始日期：20190410
	EndDate      string // 结束日期：20190415
}

// 响应结构体

type GetAnalysisTerminalInfoResponsedatasum struct {
	AccumulativeNewUserCountRatio interface{} `json:"accumulative_new_user_count_ratio"` // 新用户分布，该品牌(或设备型号、操作系统、分辨率、联网方式）的新用户数占所有品牌(或设备型号、操作系统、分辨率、联网方式）新用户之和的比例
	AccumulativeSessionCount      interface{} `json:"accumulative_session_count"`        // 启动次数，该品牌(或设备型号、操作系统、分辨率、联网方式）的启动次数
	AccumulativeSessionCountRatio interface{} `json:"accumulative_session_count_ratio"`  // 启动次数比例分布，该品牌(或设备型号、操作系统、分辨率、联网方式）的启动次数占所有品牌(或设备型号、操作系统、分辨率、联网方式）启动次数之和的比例
	AccumulativeUserCountRatio    interface{} `json:"accumulative_user_count_ratio"`     // 启动用户分布，该品牌(或设备型号、操作系统、分辨率、联网方式）的启动用户数占所有品牌(或设备型号、操作系统、分辨率、联网方式）启动用户之和（不去重）的比例
	AverageUseTime                interface{} `json:"average_use_time"`                  // 次均使用时长，平均每一次启动小程序的时间，等于总时长/总启动次数
}

type GetAnalysisTerminalInfoResponsedatadataListItem struct {
	AccumulativeNewUserCountRatio interface{} `json:"accumulative_new_user_count_ratio"` // 新用户分布，该品牌(或设备型号、操作系统、分辨率、联网方式）的新用户数占所有品牌(或设备型号、操作系统、分辨率、联网方式）新用户之和的比例
	AccumulativeSessionCount      interface{} `json:"accumulative_session_count"`        // 启动次数，该品牌(或设备型号、操作系统、分辨率、联网方式）的启动次数
	AccumulativeSessionCountRatio interface{} `json:"accumulative_session_count_ratio"`  // 启动次数比例分布，该品牌(或设备型号、操作系统、分辨率、联网方式）的启动次数占所有品牌(或设备型号、操作系统、分辨率、联网方式）启动次数之和的比例
	AccumulativeUserCountRatio    interface{} `json:"accumulative_user_count_ratio"`     // 启动用户分布，该品牌(或设备型号、操作系统、分辨率、联网方式）的启动用户数占所有品牌(或设备型号、操作系统、分辨率、联网方式）启动用户之和（不去重）的比例
	AverageUseTime                interface{} `json:"average_use_time"`                  // 次均使用时长，平均每一次启动小程序的时间，等于总时长/总启动次数
	Name                          interface{} `json:"name"`                              //
}

type GetAnalysisTerminalInfoResponsedata struct {
	DataList []GetAnalysisTerminalInfoResponsedatadataListItem `json:"data_list"` // 数据列表(指标内容)
	Offset   int64                                             `json:"offset"`    // 数据偏移位置
	Sum      GetAnalysisTerminalInfoResponsedatasum            `json:"sum"`       // 数据综合(指标内容)
	TimeSpan string                                            `json:"time_span"` // 时间范围
	Total    int64                                             `json:"total"`     // 数据量
}

type GetAnalysisTerminalInfoResponse struct {
	Data      GetAnalysisTerminalInfoResponsedata `json:"data"`       // 响应参数
	Errno     int64                               `json:"errno"`      // 状态码
	ErrMsg    string                              `json:"msg"`        // 错误信息
	ErrorCode int64                               `json:"error_code"` // openapi 错误码
	ErrorMsg  string                              `json:"error_msg"`  // openapi 错误信息
}

// GetAnalysisTerminalInfo
func GetAnalysisTerminalInfo(params *GetAnalysisTerminalInfoRequest) (*GetAnalysisTerminalInfoResponsedata, error) {
	var (
		err        error
		defaultRet *GetAnalysisTerminalInfoResponsedata
	)
	respData := &GetAnalysisTerminalInfoResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/data/getanalysisterminal")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("terminal_type", params.TerminalType)
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
