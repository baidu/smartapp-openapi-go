package pay

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetOtherDetailRequest 请求结构体
type GetOtherDetailRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	CurrentPage int64  // 当前页数。起始为 1
	PageSize    int64  // 分页。每页数量
	StartTime   string // 开始日期（格式如： 2019-01-02 ）
	EndTime     string // 结束日期（格式如： 2019-01-02 ）
}

// 响应结构体

type GetOtherDetailResponsedatadataItem struct {
	OperateTime  string `json:"operate_time"`   // 操作时间（格式为 yyyy-MM-dd HH:mm:ss ，如： 2019-05-21 23:59:59 ）
	OptType      string `json:"opt_type"`       // 操作类型（包括：成交、使用、取消使用、退款 ）
	OrderID      int64  `json:"order_id"`       // 平台订单号（百度平台订单ID ）
	RefMoney     string `json:"ref_money"`      // 金额（元）
	ThirdOrderID string `json:"third_order_id"` // 第三方订单号
}

type GetOtherDetailResponsedata struct {
	Data                []GetOtherDetailResponsedatadataItem `json:"data"`                  // 收入列表详情数据
	OrderAmount         string                               `json:"order_amount"`          // 货款总额（元）
	RuleType            string                               `json:"rule_type"`             // 规则类型。收费规则，目前只有佣金模式
	SettlementEndTime   string                               `json:"settlement_end_time"`   // 账单周期结束时间（格式为 yyyy-MM-dd HH:mm:ss ，如： 2019-05-21 23:59:59 ）
	SettlementStartTime string                               `json:"settlement_start_time"` // 账单周期开始时间（格式为 yyyy-MM-dd HH:mm:ss ，如： 2019-05-21 23:59:59 ）
	TotalAmount         string                               `json:"total_amount"`          // 入账金额（元）
}

type GetOtherDetailResponse struct {
	Data      GetOtherDetailResponsedata `json:"data"`       //
	Errno     int64                      `json:"errno"`      //
	ErrMsg    string                     `json:"msg"`        // 错误信息
	ErrorCode int64                      `json:"error_code"` // openapi 错误码
	ErrorMsg  string                     `json:"error_msg"`  // openapi 错误信息
}

// GetOtherDetail
func GetOtherDetail(params *GetOtherDetailRequest) (*GetOtherDetailResponsedata, error) {
	var (
		err        error
		defaultRet *GetOtherDetailResponsedata
	)
	respData := &GetOtherDetailResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/financebalance/otherdetail")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("current_page", fmt.Sprintf("%v", params.CurrentPage))
	client.AddGetParam("page_size", fmt.Sprintf("%v", params.PageSize))
	client.AddGetParam("start_time", params.StartTime)
	client.AddGetParam("end_time", params.EndTime)
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
