package pay

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// UpdatePaymentServiceRequest 请求结构体
type UpdatePaymentServiceRequest struct {
	AccessToken    string // 授权小程序的接口调用凭据
	AppName        string // 服务名称。
	ServicePhone   string // 服务电话。
	PoolCashPledge int64  // 打款预留。提现后的保留金额
}

// 响应结构体

type UpdatePaymentServiceResponsedata struct {
}

type UpdatePaymentServiceResponse struct {
	Data      UpdatePaymentServiceResponsedata `json:"data"`       // 响应参数
	Errno     int64                            `json:"errno"`      // 状态码
	ErrMsg    string                           `json:"msg"`        // 错误信息
	ErrorCode int64                            `json:"error_code"` // openapi 错误码
	ErrorMsg  string                           `json:"error_msg"`  // openapi 错误信息
}

// UpdatePaymentService
func UpdatePaymentService(params *UpdatePaymentServiceRequest) (*UpdatePaymentServiceResponsedata, error) {
	var (
		err        error
		defaultRet *UpdatePaymentServiceResponsedata
	)
	respData := &UpdatePaymentServiceResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/update")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("app_name", params.AppName)
	client.AddPostParam("service_phone", params.ServicePhone)
	client.AddPostParam("pool_cash_pledge", fmt.Sprintf("%v", params.PoolCashPledge))

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
