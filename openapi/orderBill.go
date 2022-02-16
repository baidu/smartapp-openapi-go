package openapi


// OrderBillRequest 请求结构体
type OrderBillRequest struct {
	AccessToken string // 接口调用凭证
	BillTime    string // 对账单日期格式 yyyy-MM-dd
	PmAppKey    string // 调起百度收银台的支付服务 appKey
}

// 响应结构体

type OrderBillResponsedatadataItem struct {
	CreateTime   string `json:"createTime"`   // 创建时间
	DownloadName string `json:"downloadName"` // 账单名称
	ExportStatus int64  `json:"exportStatus"` // 导出进度
	URL          string `json:"url"`          // 下载地址
}

type OrderBillResponsedata struct {
	Data       []OrderBillResponsedatadataItem `json:"data"`       // 响应对象
	TotalCount int64                           `json:"totalCount"` // 总数
}

type OrderBillResponse struct {
	Data      OrderBillResponsedata `json:"data"`       // 响应对象
	Errno     int64                 `json:"errno"`      // 错误码
	ErrMsg    string                `json:"msg"`        // 错误信息
	ErrorCode int64                 `json:"error_code"` // openapi 错误码
	ErrorMsg  string                `json:"error_msg"`  // openapi 错误信息
}

// OrderBill
func OrderBill(params *OrderBillRequest) (*OrderBillResponsedata, error) {
	var (
		err        error
		defaultRet *OrderBillResponsedata
	)
	respData := &OrderBillResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/orderBill")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("billTime", params.BillTime)
	client.AddGetParam("pmAppKey", params.PmAppKey)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)

	err = client.Do()
	if err != nil {
		return defaultRet, err
	}
	err = client.Convert(respData)
	if err != nil {
		return defaultRet, err
	}
	if respData.ErrorCode != 0 {
		return defaultRet, &OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}
	if respData.Errno != 0 {
		return defaultRet, &APIError{respData.Errno, respData.ErrMsg, respData}
	}

	return &respData.Data, nil
}
