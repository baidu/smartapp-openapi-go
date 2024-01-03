package openapi

// FindByTpOrderIDRequest 请求结构体
type FindByTpOrderIDRequest struct {
	AccessToken string // 接口调用凭证
	TpOrderID   string // 开发者订单 ID
	PmAppKey    string // 调起百度收银台的支付服务 appKey
}

// 响应结构体

type FindByTpOrderIDResponsedata struct {
	AppID         int64  `json:"appId"`         // 小程序appid
	AppKey        string `json:"appKey"`        // 小程序appkey
	BizInfo       string `json:"bizInfo"`       // 业务扩展字段
	Count         int64  `json:"count"`         // 数量
	CreateTime    int64  `json:"createTime"`    // 创建时间
	DealID        int64  `json:"dealId"`        // 跳转百度收银台支付必带参数之一
	OpenID        string `json:"openId"`        // 小程序用户id
	OrderID       int64  `json:"orderId"`       // 百度订单 ID
	OriPrice      int64  `json:"oriPrice"`      // 原价
	ParentOrderID int64  `json:"parentOrderId"` // 购物车订单父订单ID
	ParentType    int64  `json:"parentType"`    // 订单类型
	PayMoney      int64  `json:"payMoney"`      // 支付金额
	SettlePrice   int64  `json:"settlePrice"`   // 结算金额
	Status        int64  `json:"status"`        // 订单状态 1未支付 2已支付
	SubStatus     int64  `json:"subStatus"`     // 订单子状态
	TotalMoney    int64  `json:"totalMoney"`    // 总金额
	TpID          int64  `json:"tpId"`          // tpid
	TpOrderID     string `json:"tpOrderId"`     // 开发者订单ID
	TradeNo       string `json:"tradeNo"`       // 支付单号
	Type          int64  `json:"type"`          // ordertype
	UserID        int64  `json:"userId"`        // 用户 id 与支付状态通知中的保持一致
}

type FindByTpOrderIDResponse struct {
	Data      FindByTpOrderIDResponsedata `json:"data"`       // 响应对象
	Errno     int64                       `json:"errno"`      // 错误码
	ErrMsg    string                      `json:"msg"`        // 错误信息
	ErrorCode int64                       `json:"error_code"` // openapi 错误码
	ErrorMsg  string                      `json:"error_msg"`  // openapi 错误信息
}

// FindByTpOrderID
func FindByTpOrderID(params *FindByTpOrderIDRequest) (*FindByTpOrderIDResponsedata, error) {
	var (
		err        error
		defaultRet *FindByTpOrderIDResponsedata
	)
	respData := &FindByTpOrderIDResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/findByTpOrderId")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("tpOrderId", params.TpOrderID)
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
