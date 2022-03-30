package pay

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// UpdateBindPaymentServiceRequest 请求结构体
type UpdateBindPaymentServiceRequest struct {
	AccessToken        string // 授权小程序的接口调用凭据
	AppName            string // 服务名称。支付服务的名称
	ServicePhone       string // 服务电话。
	BankAccount        string // 银行账户。银行卡开户名
	BankCard           string // 银行卡号。
	BankName           string // 所属银行。由数据字典接口取
	BankBranch         string // 支行信息。自由输入，例如：招商银行北京上地支行
	OpenProvince       string // 开户省份。由数据字典接口取
	OpenCity           string // 开户城市。由数据字典接口取
	PaymentDays        int64  // 结算周期。由数据字典接口取
	CommissionRate     int64  // 佣金比例。固定传 6，小程序固定为千分之六(6)
	PoolCashPledge     int64  // 打款预留（元）。提现后的保留金额
	DayMaxFrozenAmount int64  // 每日退款上限(元)。每天最大退款限额10000元
	BankPhoneNumber    string // 银行卡预留手机号
}

// 响应结构体

type UpdateBindPaymentServiceResponsedata struct {
}

type UpdateBindPaymentServiceResponse struct {
	Data      UpdateBindPaymentServiceResponsedata `json:"data"`       // 响应参数
	Errno     int64                                `json:"errno"`      // 状态码
	ErrMsg    string                               `json:"msg"`        // 错误信息
	ErrorCode int64                                `json:"error_code"` // openapi 错误码
	ErrorMsg  string                               `json:"error_msg"`  // openapi 错误信息
}

// UpdateBindPaymentService
func UpdateBindPaymentService(params *UpdateBindPaymentServiceRequest) (*UpdateBindPaymentServiceResponsedata, error) {
	var (
		err        error
		defaultRet *UpdateBindPaymentServiceResponsedata
	)
	respData := &UpdateBindPaymentServiceResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/updatebindservice")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("app_name", params.AppName)
	client.AddPostParam("service_phone", params.ServicePhone)
	client.AddPostParam("bank_account", params.BankAccount)
	client.AddPostParam("bank_card", params.BankCard)
	client.AddPostParam("bank_name", params.BankName)
	client.AddPostParam("bank_branch", params.BankBranch)
	client.AddPostParam("open_province", params.OpenProvince)
	client.AddPostParam("open_city", params.OpenCity)
	client.AddPostParam("payment_days", fmt.Sprintf("%v", params.PaymentDays))
	client.AddPostParam("commission_rate", fmt.Sprintf("%v", params.CommissionRate))
	client.AddPostParam("pool_cash_pledge", fmt.Sprintf("%v", params.PoolCashPledge))
	client.AddPostParam("day_max_frozen_amount", fmt.Sprintf("%v", params.DayMaxFrozenAmount))
	client.AddPostParam("bank_phone_number", params.BankPhoneNumber)

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
