package pay

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetExpenseDetailRequest 请求结构体
type GetExpenseDetailRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	TaskID      int64  // 支付凭证 id。
}

// 响应结构体

type GetExpenseDetailResponsedatadataItem struct {
	PaidAccount   string `json:"paid_account"`   // 收款账号。收款的银行卡号
	PaidMoney     string `json:"paid_money"`     // 交易收款金额（元）
	PaymentStatus string `json:"payment_status"` // 收款状态。状态包含：待出纳支付、打款中、成功、失败
	ReceiveUser   string `json:"receive_user"`   // 收款人姓名
	Remark        string `json:"remark"`         // 备注信息。交易支出备注信息
}

type GetExpenseDetailResponsedata struct {
	Data []GetExpenseDetailResponsedatadataItem `json:"data"` // 数据详情
}

type GetExpenseDetailResponse struct {
	Data      GetExpenseDetailResponsedata `json:"data"`       //
	Errno     int64                        `json:"errno"`      //
	ErrMsg    string                       `json:"msg"`        // 错误信息
	ErrorCode int64                        `json:"error_code"` // openapi 错误码
	ErrorMsg  string                       `json:"error_msg"`  // openapi 错误信息
}

// GetExpenseDetail
func GetExpenseDetail(params *GetExpenseDetailRequest) (*GetExpenseDetailResponsedata, error) {
	var (
		err        error
		defaultRet *GetExpenseDetailResponsedata
	)
	respData := &GetExpenseDetailResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/financeexpense/paydetail")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("task_id", fmt.Sprintf("%v", params.TaskID))
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
