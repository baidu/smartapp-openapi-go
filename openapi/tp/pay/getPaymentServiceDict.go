package pay

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetPaymentServiceDictRequest 请求结构体
type GetPaymentServiceDictRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
}

// 响应结构体

type GetPaymentServiceDictResponseprovinceListItemcitiesItem struct {
	CityID   int64  `json:"cityId"`   //
	CityName string `json:"cityName"` //
}

type GetPaymentServiceDictResponseprovinceListItem struct {
	Cities       []GetPaymentServiceDictResponseprovinceListItemcitiesItem `json:"cities"`       //
	ProvinceID   int64                                                     `json:"provinceId"`   //
	ProvinceName string                                                    `json:"provinceName"` //
}

type GetPaymentServiceDictResponsebankListItembanksItem struct {
	BankID   int64  `json:"bankId"`   //
	BankName string `json:"bankName"` //
	LogoURL  string `json:"logoUrl"`  //
}

type GetPaymentServiceDictResponsebankListItem struct {
	Banks     []GetPaymentServiceDictResponsebankListItembanksItem `json:"banks"`     //
	Character string                                               `json:"character"` //
}

type GetPaymentServiceDictResponse struct {
	BankList       []GetPaymentServiceDictResponsebankListItem     `json:"bank_list"`      // 开户银行 （支付服务 bank_name）
	CommissionRate map[string]string                               `json:"commissionRate"` // 佣金比例，小程序固定为千分之六（支付服务 commission_rate）
	PaymentDays    map[string]string                               `json:"paymentDays"`    // 结算周期（支付服务 payment_days）
	ProvinceList   []GetPaymentServiceDictResponseprovinceListItem `json:"province_list"`  // 开户省市（支付服务 open_province，open_city ）
	ErrorCode      int64                                           `json:"error_code"`     // openapi 错误码
	ErrorMsg       string                                          `json:"error_msg"`      // openapi 错误信息
	Data           GetPaymentServiceDictResponsedata
}

type GetPaymentServiceDictResponsedata struct {
	BankList       []GetPaymentServiceDictResponsebankListItem     `json:"bank_list"`      // 开户银行 （支付服务 bank_name）
	CommissionRate map[string]string                               `json:"commissionRate"` // 佣金比例，小程序固定为千分之六（支付服务 commission_rate）
	PaymentDays    map[string]string                               `json:"paymentDays"`    // 结算周期（支付服务 payment_days）
	ProvinceList   []GetPaymentServiceDictResponseprovinceListItem `json:"province_list"`  // 开户省市（支付服务 open_province，open_city ）
}

// GetPaymentServiceDict
func GetPaymentServiceDict(params *GetPaymentServiceDictRequest) (*GetPaymentServiceDictResponsedata, error) {
	var (
		err        error
		defaultRet *GetPaymentServiceDictResponsedata
	)
	respData := &GetPaymentServiceDictResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/dict")
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

	return &respData.Data, nil
}
