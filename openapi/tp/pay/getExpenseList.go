package pay

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetExpenseListRequest 请求结构体
type GetExpenseListRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	CurrentPage int64  // 当前页数。起始为 1
	PageSize    int64  // 分页。每页数量
	StartTime   string // 开始日期（格式如： 2019-01-02 ）
	EndTime     string // 结束日期（格式如： 2019-01-02 ）
}

// 响应结构体

type GetExpenseListResponsedatadataItem struct {
	EndTime     string `json:"end_time"`     // 账期结束时间。格式为 yyyy-MM-dd HH:mm:ss，如：2019-06-07 00:00:00
	Money       string `json:"money"`        // 金额（元）。当前交易的金额
	OperateTime string `json:"operate_time"` // 日期。格式为 yyyy-MM-dd HH:mm:ss，如：2019-06-07 00:00:00
	PayStatus   string `json:"pay_status"`   // 支出打款状态。状态包含：待出纳支付、打款中、成功、失败
	StartTime   string `json:"start_time"`   // 账期开始时间。格式为 yyyy-MM-dd HH:mm:ss，如：2019-06-07 00:00:00
	TaskID      int64  `json:"task_id"`      // 支付凭证 id。用于查询支出打款详情
	TypeID      string `json:"type_id"`      // 类型。如：收入货款、调整款、自动打款等
}

type GetExpenseListResponsedata struct {
	Data         []GetExpenseListResponsedatadataItem `json:"data"`          // 支出列表详情
	ExpenseCount int64                                `json:"expense_count"` // 支出次数总和
	ExpenseMoney string                               `json:"expense_money"` // 支出金额（元）。所有支出次数金额的汇总
	TotalCount   int64                                `json:"total_count"`   // 数据总条数
}

type GetExpenseListResponse struct {
	Data      GetExpenseListResponsedata `json:"data"`       //
	Errno     int64                      `json:"errno"`      //
	ErrMsg    string                     `json:"msg"`        // 错误信息
	ErrorCode int64                      `json:"error_code"` // openapi 错误码
	ErrorMsg  string                     `json:"error_msg"`  // openapi 错误信息
}

// GetExpenseList
func GetExpenseList(params *GetExpenseListRequest) (*GetExpenseListResponsedata, error) {
	var (
		err        error
		defaultRet *GetExpenseListResponsedata
	)
	respData := &GetExpenseListResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/financeexpense")
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
