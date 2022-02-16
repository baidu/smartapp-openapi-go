package openapi


// BatchGetCouponRequest 请求结构体
type BatchGetCouponRequest struct {
	AccessToken string // 接口调用凭证
}

// 响应结构体

type BatchGetCouponResponsedatadataListItembaseInfodateInfo struct {
	BeginTimestamp    int64 `json:"beginTimestamp"`    // 使用开始时间，当 type 为 1 时，beginTimestamp 必传且生效；
	EndTimestamp      int64 `json:"endTimestamp"`      // 使用结束时间，当 type 为 1 时，endTimestamp 必传且生效；
	GetEndTimestamp   int64 `json:"getEndTimestamp"`   // 结束领取时间
	GetStartTimestamp int64 `json:"getStartTimestamp"` // 开始领取时间
	TimeUnit          int64 `json:"timeUnit"`          // 时间单位：1-时；2-天；3-月；当 type 为 2 时，timeUnit 必传且生效；
	TimeValue         int64 `json:"timeValue"`         // 时间值；当 type 为 2 时，timeValue 必传且生效；
	Type              int64 `json:"type"`              // 券使用时间类型： 1：开发者设置使用开始和结束时间，此时，beginTimestamp 和 endTimestamp 必传。 2：领取之后，多久可使用，此时，timeUnit 和 timeValue 必传。相对时间：当规定领取 5 日后失效，10 月 1 日 23:00 领取后，10 月 6 日 23:00 失效。 4：领取之后，多久（自然日）失效，此时，timeUnit 和 timeValue 必传，timeUnit 只能设置为天（timeUnit=2）。自然相对时间：当规定领取 5 个自然日后失效，10 月 1 日 23:00 领取后，10 月 6 日 00:00 失效。
}

type BatchGetCouponResponsedatadataListItembaseInfo struct {
	Color    string                                                 `json:"color"`    // 卡券背景色
	DateInfo BatchGetCouponResponsedatadataListItembaseInfodateInfo `json:"dateInfo"` // 使用日期，有效期的信息
	GetLimit int64                                                  `json:"getLimit"` // 每人领取次数限制
	Title    string                                                 `json:"title"`    // 优惠券名称
}

type BatchGetCouponResponsedatadataListItem struct {
	BaseInfo    BatchGetCouponResponsedatadataListItembaseInfo `json:"baseInfo"`    // 基本的卡券数据，所有卡券通用。
	CallbackURL string                                         `json:"callbackUrl"` // 卡券领取事件推送地址
	CouponType  string                                         `json:"couponType"`  // 卡券类型
	Description string                                         `json:"description"` // 使用须知：卡券使用方法的介绍
	Discount    int64                                          `json:"discount"`    // 折扣券专用，表示打折力度（格式为百分比），填 80 就是八折。
}

type BatchGetCouponResponsedata struct {
	DataList []BatchGetCouponResponsedatadataListItem `json:"dataList"` // 卡券信息数组
	PageNo   int64                                    `json:"pageNo"`   // 页码
	Total    int64                                    `json:"total"`    // 总数量
}

type BatchGetCouponResponse struct {
	Data      BatchGetCouponResponsedata `json:"data"`       // 响应对象
	Errno     int64                      `json:"errno"`      // 错误码
	ErrMsg    string                     `json:"msg"`        // 错误信息
	ErrorCode int64                      `json:"error_code"` // openapi 错误码
	ErrorMsg  string                     `json:"error_msg"`  // openapi 错误信息
}

// BatchGetCoupon
func BatchGetCoupon(params *BatchGetCouponRequest) (*BatchGetCouponResponsedata, error) {
	var (
		err        error
		defaultRet *BatchGetCouponResponsedata
	)
	respData := &BatchGetCouponResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/v1.0/coupon/batchGet")
	client.AddGetParam("access_token", params.AccessToken)
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
