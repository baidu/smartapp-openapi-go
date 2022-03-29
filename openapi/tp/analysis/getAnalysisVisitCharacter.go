package analysis

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAnalysisVisitCharacterRequest 请求结构体
type GetAnalysisVisitCharacterRequest struct {
	AccessToken   string // 授权小程序的接口调用凭据
	CharacterType string // 习惯分类：depth/time/interval/frequency
	StartIndex    int64  // 数据偏移位置，接口默认返回 20 条数据，可使用该偏移量进行翻页查看
	StartDate     string // 开始日期：20190410
	EndDate       string // 结束日期：20190415
}

// 响应结构体

type GetAnalysisVisitCharacterResponsedatasum struct {
	SessionCount      interface{} `json:"session_count"`       // 启动用户数，启动过小程序的用户（多次启动不重复计）
	SessionCountRatio interface{} `json:"session_count_ratio"` // 启动用户数量比例，当前类型启动过小程序的用户在所有类型启动过小程序用户的比率。
	UserCount         interface{} `json:"user_count"`          // 启动用户数 (character_type为frequency返回的指标)，启动过小程序的用户（多次启动不重复计）
	UserCountRatio    interface{} `json:"user_count_ratio"`    // 启动用户分布(character_type为frequency返回的指标)，当前类型启动过小程序的用户在所有类型启动过小程序用户的比率。
}

type GetAnalysisVisitCharacterResponsedata struct {
	DataList []map[string]interface{}                 `json:"data_list"` // 数据列表(指标内容)
	Offset   int64                                    `json:"offset"`    // 数据偏移位置
	Sum      GetAnalysisVisitCharacterResponsedatasum `json:"sum"`       // 数据综合(指标内容)
	TimeSpan string                                   `json:"time_span"` // 时间范围
	Total    int64                                    `json:"total"`     // 数据量
}

type GetAnalysisVisitCharacterResponse struct {
	Data      GetAnalysisVisitCharacterResponsedata `json:"data"`       // 响应参数
	Errno     int64                                 `json:"errno"`      // 状态码
	ErrMsg    string                                `json:"msg"`        // 错误信息
	ErrorCode int64                                 `json:"error_code"` // openapi 错误码
	ErrorMsg  string                                `json:"error_msg"`  // openapi 错误信息
}

// GetAnalysisVisitCharacter
func GetAnalysisVisitCharacter(params *GetAnalysisVisitCharacterRequest) (*GetAnalysisVisitCharacterResponsedata, error) {
	var (
		err        error
		defaultRet *GetAnalysisVisitCharacterResponsedata
	)
	respData := &GetAnalysisVisitCharacterResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/data/getanalysisvisitcharacter")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("character_type", params.CharacterType)
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
