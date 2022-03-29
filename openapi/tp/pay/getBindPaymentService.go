package pay

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetBindPaymentServiceRequest 请求结构体
type GetBindPaymentServiceRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
}

// 响应结构体

type GetBindPaymentServiceResponsedata struct {
	AppName            string `json:"app_name"`              // 服务名称。支付服务的名称
	BankAccount        string `json:"bank_account"`          // 银行账户。银行卡开户名
	BankBranch         string `json:"bank_branch"`           // 支行信息。自由输入，例如：招商银行北京上地支行
	BankCard           string `json:"bank_card"`             // 银行卡号。
	BankName           string `json:"bank_name"`             // 所属银行。由数据字典接口取
	CommissionRate     int64  `json:"commission_rate"`       // 佣金比例。固定传 6，小程序固定为千分之六(6)
	DayMaxFrozenAmount int64  `json:"day_max_frozen_amount"` // 每日退款上限(元)。每天最大退款限额10000元
	DealID             string `json:"deal_id"`               // 百度收银台的财务结算凭证。详见电商技术平台术语
	FailReason         string `json:"fail_reason"`           // 驳回回执。
	OpenCity           string `json:"open_city"`             // 开户城市。由数据字典接口取
	OpenProvince       string `json:"open_province"`         // 开户省份。由数据字典接口取
	OpenStatus         int64  `json:"open_status"`           // 开通状态. 0:新建 1:审核中 2:审核通过 3:驳回
	PaymentDays        int64  `json:"payment_days"`          // 结算周期。由数据字典接口取
	PlatformPublicKey  string `json:"platform_public_key"`   // 平台公钥。详见电商技术平台术语
	PmAppID            int64  `json:"pm_app_id"`             // 服务id。支付服务内部标示 id
	PmAppKey           string `json:"pm_app_key"`            // 服务Key。支付服务内部标示key
	PoolCashPledge     int64  `json:"pool_cash_pledge"`      // 打款预留（元）。提现后的保留金额
	ServicePhone       string `json:"service_phone"`         // 服务电话。
}

type GetBindPaymentServiceResponse struct {
	Data      GetBindPaymentServiceResponsedata `json:"data"`       // 响应参数
	Errno     int64                             `json:"errno"`      // 状态码
	ErrMsg    string                            `json:"msg"`        // 错误信息
	ErrorCode int64                             `json:"error_code"` // openapi 错误码
	ErrorMsg  string                            `json:"error_msg"`  // openapi 错误信息
}

// GetBindPaymentService
func GetBindPaymentService(params *GetBindPaymentServiceRequest) (*GetBindPaymentServiceResponsedata, error) {
	var (
		err        error
		defaultRet *GetBindPaymentServiceResponsedata
	)
	respData := &GetBindPaymentServiceResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/getbindservice")
	client.AddGetParam("access_token", params.AccessToken)
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
