package pay

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetIncomeDetailRequest 请求结构体
type GetIncomeDetailRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	StartTime   string // 开始日期。格式如：2019-01-02。
	CurrentPage int64  // 当前页数。起始为 1
	PageSize    int64  // 分页。每页数量
}

// 响应结构体

type GetIncomeDetailResponsedatadataItem struct {
	Money        string `json:"money"`          // 金额（元）
	OperateTime  string `json:"operate_time"`   // 操作时间。格式为 yyyy-MM-dd，如：2019-06-07
	OptType      string `json:"opt_type"`       // 操作类型。包括：成交、使用、取消使用、退款
	OrderID      int64  `json:"order_id"`       // 平台订单号。百度平台订单ID
	ThirdOrderID string `json:"third_order_id"` // 第三方订单号
}

type GetIncomeDetailResponsedata struct {
	CurrentPage int64 `json:"current_page"` // 当前页数

	Data         []GetIncomeDetailResponsedatadataItem `json:"data"`          // 列表详情数据
	Income       string                                `json:"income"`        // 货款（元）
	IncomeAmount string                                `json:"income_amount"` // 验证(成交)总额（元）
	OperateTime  string                                `json:"operate_time"`  // 操作日期。格式为 yyyy-MM-dd HH:mm:ss，如：2019-05-21 23:59:59
	PageSize     int64                                 `json:"page_size"`     // 分页每页数量
	Quantity     int64                                 `json:"quantity"`      // 验证（成交）数量
	RefundAmount string                                `json:"refund_amount"` // 退款总额（元）
	TotalCount   int64                                 `json:"total_count"`   // 总数据数量
}

type GetIncomeDetailResponse struct {
	Data      GetIncomeDetailResponsedata `json:"data"`       //
	Errno     int64                       `json:"errno"`      //
	ErrMsg    string                      `json:"msg"`        // 错误信息
	ErrorCode int64                       `json:"error_code"` // openapi 错误码
	ErrorMsg  string                      `json:"error_msg"`  // openapi 错误信息
}

// GetIncomeDetail
func GetIncomeDetail(params *GetIncomeDetailRequest) (*GetIncomeDetailResponsedata, error) {
	var (
		err        error
		defaultRet *GetIncomeDetailResponsedata
	)
	respData := &GetIncomeDetailResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/financebalance/incomedetail")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("start_time", params.StartTime)
	client.AddGetParam("current_page", fmt.Sprintf("%v", params.CurrentPage))
	client.AddGetParam("page_size", fmt.Sprintf("%v", params.PageSize))
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
