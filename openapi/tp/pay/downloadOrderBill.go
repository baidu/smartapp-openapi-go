package pay

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// DownloadOrderBillRequest 请求结构体
type DownloadOrderBillRequest struct {
	AccessToken string // 第三方平台的接口调用凭据
	BillTime    string // 对账单日期格式 yyyy-MM-dd
	PmAppKey    string // 调起百度收银台的支付服务 appKey
}

// 响应结构体

type DownloadOrderBillResponsedatadataItem struct {
	CreateTime   string  `json:"createTime"`   //  创建时间
	DownloadName string  `json:"downloadName"` // 账单名称
	ExportStatus float64 `json:"exportStatus"` // 导出进度
	URL          string  `json:"url"`          // 下载地址
}

type DownloadOrderBillResponsedata struct {
	Data       []DownloadOrderBillResponsedatadataItem `json:"data"`       //
	TotalCount int64                                   `json:"totalCount"` //
}

type DownloadOrderBillResponse struct {
	Data      DownloadOrderBillResponsedata `json:"data"`       // 响应参数
	Errno     int64                         `json:"errno"`      // 状态码
	ErrMsg    string                        `json:"msg"`        // 错误信息
	ErrorCode int64                         `json:"error_code"` // openapi 错误码
	ErrorMsg  string                        `json:"error_msg"`  // openapi 错误信息
}

// DownloadOrderBill
func DownloadOrderBill(params *DownloadOrderBillRequest) (*DownloadOrderBillResponsedata, error) {
	var (
		err        error
		defaultRet *DownloadOrderBillResponsedata
	)
	respData := &DownloadOrderBillResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/tp/orderBill")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("billTime", params.BillTime)
	client.AddGetParam("pmAppKey", params.PmAppKey)
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
