package pay

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetFinanceBalanceRequest 请求结构体
type GetFinanceBalanceRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	StartTime   string // 开始日期。格式如：2019-01-02。
	EndTime     string // 结束日期。格式如： 2019-01-02。
	CurrentPage int64  // 当前页数。起始为 1
	PageSize    int64  // 分页。每页数量
}

// 响应结构体

type GetFinanceBalanceResponsedatadataItem struct {
	AdjustAmount string `json:"adjust_amount"` // 调整款（元）。如：商家作弊，交易会扣钱，会以调整款的形式。
	Income       string `json:"income"`        // 货款金额（元）
	OperateTime  string `json:"operate_time"`  // 日期。格式为 yyyy-MM-dd，如：2019-06-07
	Others       string `json:"others"`        // 其他款项。佣金、返点与营销费用等金额
}

type GetFinanceBalanceResponsedata struct {
	Adjustment  string `json:"adjustment"`   // 调整款（元）。如：商家作弊，交易会扣钱，会以调整款的形式。
	CurrentPage int64  `json:"current_page"` // 当前页码

	Data        []GetFinanceBalanceResponsedatadataItem `json:"data"`         // 收入列表详情数据。
	FreeBalance string                                  `json:"free_balance"` // 账户余额（元）。
	Income      string                                  `json:"income"`       // 货款总金额（元）。
	Others      string                                  `json:"others"`       // 其它款项（元）。一般为佣金、返点与营销费用等
	PageSize    int64                                   `json:"page_size"`    // 分页每页数量
	PaymentDue  string                                  `json:"payment_due"`  // 支付时间。格式为 yyyy-MM-dd HH:mm:ss，如：2019-05-21 23:59:59
	Period      int64                                   `json:"period"`       // 当前付款周期（天）。
	TotalCount  int64                                   `json:"total_count"`  // 总页码数
}

type GetFinanceBalanceResponse struct {
	Data      GetFinanceBalanceResponsedata `json:"data"`       //
	Errno     int64                         `json:"errno"`      //
	ErrMsg    string                        `json:"msg"`        // 错误信息
	ErrorCode int64                         `json:"error_code"` // openapi 错误码
	ErrorMsg  string                        `json:"error_msg"`  // openapi 错误信息
}

// GetFinanceBalance
func GetFinanceBalance(params *GetFinanceBalanceRequest) (*GetFinanceBalanceResponsedata, error) {
	var (
		err        error
		defaultRet *GetFinanceBalanceResponsedata
	)
	respData := &GetFinanceBalanceResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/financebalance")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("start_time", params.StartTime)
	client.AddGetParam("end_time", params.EndTime)
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
