package openapi

import (
	"encoding/json"
)

// SubmitSkuCouponRequest 请求结构体
type SubmitSkuCouponRequest struct {
	AccessToken string // 接口调用凭证
	PostBody    []SubmitSkuCouponRequestBody
}

type SubmitSkuCouponRequestpriceInfo struct {
	OrgPrice  string `json:"org_price"`  // 付费优惠券：优惠活动前的服务原价格，注意，是以分为单位
	RealPrice string `json:"real_price"` // 付费优惠券：当前服务的实际成交价格，注意，是以分为单位
}

type SubmitSkuCouponRequestBody struct {
	Desc      string                          `json:"desc"`       // 优惠券简介，优惠券的文字解释说明，8~34 个字符(汉字占 2 字符)
	Images    []string                        `json:"images"`     // 优惠券图片比例为 1：1，像素不得低于 576 * 576,支持 png、jpg 图片内容要求：图片清晰、干净，不要出现令人不适的内容；不能出现严重影响用户理解的内容截断问题；图片无水印、二维码相关性&一致性：图片与标题、优惠落地页内容相关、信息一致
	Path      string                          `json:"path"`       // 智能小程序落地页链接，免费优惠券 path 填写格式为 /pages/detail/highVersionIndex/?biz_id=2&biz_app_id= 小程序 appKey&coupon_template_id= 卡券 id&is_activity=0
	PriceInfo SubmitSkuCouponRequestpriceInfo `json:"price_info"` // 服务价格，详见：price_info 字段说明
	Region    string                          `json:"region"`     // 服务地域，参考附录二，省市之间用英文中划线分割，多个地区之间用英文逗号分割
	Schema    string                          `json:"schema"`     // 优惠券的具体信息，详见：coupon_info
	Title     string                          `json:"title"`      // 优惠券标题：活动优惠信息说明，12-30 个字符(汉字占 2 字符)；不允许有特殊符号；优惠活动信息必须真实；需要清晰地说明商品内容，说明券的品牌（如肯德基、爱奇艺）、优惠主体（如 30 元代金券、汉堡薯条炸鸡兑换券）
	TradeType int64                           `json:"trade_type"` // 服务类目编码，参考附录一
}

// 响应结构体

type SubmitSkuCouponResponse struct {
	Data      string `json:"data"`       // 响应参数（如有提交失败的数据会返回失败的 path 和原因）
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// SubmitSkuCoupon
func SubmitSkuCoupon(params *SubmitSkuCouponRequest) (string, error) {
	var (
		err        error
		defaultRet string
	)
	respData := &SubmitSkuCouponResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeJSON).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/server/submit/skuCoupon")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	postData := make([]SubmitSkuCouponRequestBody, 0, len(params.PostBody))
	postData = append(postData, params.PostBody...)
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

	return respData.Data, nil
}
