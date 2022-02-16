package openapi

import (
	"encoding/json"
)

// CreateCouponRequest 请求结构体
type CreateCouponRequest struct {
	AccessToken string                      // 接口调用凭证
	Description string                      // 使用须知：卡券使用方法的介绍
	CallbackURL string                      // 卡券领取事件推送地址
	CouponType  string                      // 卡券类型，当以上卡券类型无法满足时，可使用通用优惠券类型
	LeastCost   int64                       // 表示可使用的门槛金额（单位：分），不传默认为 0，即无起用门槛
	ReduceCost  int64                       // 代金券专用，表示减免金额（单位：分）
	Discount    string                      // 折扣券专用，表示打折力度（格式为百分比），填 80 就是八折。
	BaseInfo    CreateCouponRequestbaseInfo // 基本的卡券数据，所有卡券通用。
}

type CreateCouponRequestbaseInfodateInfo struct {
	BeginTimestamp    int64 `json:"beginTimestamp"`    // 使用开始时间，当 type 为 1 时，beginTimestamp 必传且生效；
	EndTimestamp      int64 `json:"endTimestamp"`      // 使用结束时间，当 type 为 1 时，endTimestamp 必传且生效；
	GetEndTimestamp   int64 `json:"getEndTimestamp"`   // 结束领取时间
	GetStartTimestamp int64 `json:"getStartTimestamp"` // 开始领取时间
	TimeUnit          int64 `json:"timeUnit"`          // 时间单位：1-时；2-天；3-月；当 type 为 2 时，timeUnit 必传且生效；
	TimeValue         int64 `json:"timeValue"`         // 时间值；当 type 为 2 时，timeValue 必传且生效；
	Type              int64 `json:"type"`              // 券使用时间类型： 1：开发者设置使用开始和结束时间，此时，beginTimestamp 和 endTimestamp 必传。 2：领取之后，多久可使用，此时，timeUnit 和 timeValue 必传。相对时间：当规定领取 5 日后失效，10 月 1 日 23:00 领取后，10 月 6 日 23:00 失效。 4：领取之后，多久（自然日）失效，此时，timeUnit 和 timeValue 必传，timeUnit 只能设置为天（timeUnit=2）。自然相对时间：当规定领取 5 个自然日后失效，10 月 1 日 23:00 领取后，10 月 6 日 00:00 失效。
}
type CreateCouponRequestbaseInfo struct {
	AppRedirectPath string                              `json:"appRedirectPath"` // 已领取的卡券，从详情頁点击「立即使用」打开小程序页面地址，不传默认打开首页
	CodeType        int64                               `json:"codeType"`        // 卡券 Code 码类型，默认为 1，1：开发者自定义 code 码，当 codeType=1 时，需要通过「上传 code 码」接口导入 Code，否则影响领券；2：系统分配 Code 码，当 codeType=2 时，开发者无需上传 Code ，quantity 要求必传非 0 且生效
	Color           string                              `json:"color"`           // 卡券背景色，支持范围： [B010 ～ B160]
	DateInfo        CreateCouponRequestbaseInfodateInfo `json:"dateInfo"`        // 有效期信息
	GetLimit        int64                               `json:"getLimit"`        // 使用日期，有效期的信息。
	Quantity        int64                               `json:"quantity"`        // 卡券库存，默认为 0，当 codeType=2 时，quantity 要求必传且生效
	Title           string                              `json:"title"`           // 优惠券名称，最多 15 字
}

// 响应结构体

type CreateCouponResponsedata struct {
	CouponID string `json:"couponId"` // 卡券 ID
}

type CreateCouponResponse struct {
	Data      CreateCouponResponsedata `json:"data"`       // 响应对象
	Errno     int64                    `json:"errno"`      // 错误码
	ErrMsg    string                   `json:"msg"`        // 错误信息
	ErrorCode int64                    `json:"error_code"` // openapi 错误码
	ErrorMsg  string                   `json:"error_msg"`  // openapi 错误信息
}

// CreateCoupon
func CreateCoupon(params *CreateCouponRequest) (*CreateCouponResponsedata, error) {
	var (
		err        error
		defaultRet *CreateCouponResponsedata
	)
	respData := &CreateCouponResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeJSON).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/v1.0/coupon/create")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	postData := map[string]interface{}{
		"description": params.Description,
		"callbackUrl": params.CallbackURL,
		"couponType":  params.CouponType,
		"leastCost":   params.LeastCost,
		"reduceCost":  params.ReduceCost,
		"discount":    params.Discount,
		"baseInfo":    params.BaseInfo,
	}
	bts, err := json.Marshal(postData)
	if err != nil {
		return defaultRet, err
	}
	client.SetBody(bts)

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
