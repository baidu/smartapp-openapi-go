package openapi

import (
	"encoding/json"
)

// SubmitSkuRequest 请求结构体
type SubmitSkuRequest struct {
	AccessToken string // 接口调用凭证
	PostBody    []SubmitSkuRequestBody
}

type SubmitSkuRequestactivityInfoItem struct {
	ActivityDesc      string `json:"activity_desc"`       // 优惠活动说明，若优惠活动类型已填，则必填
	ActivityEndTime   int64  `json:"activity_end_time"`   // 优惠结束时间（时间戳到秒），距离开始时间不长于 6 个月；若优惠活动说明已填，则必填
	ActivityPath      string `json:"activity_path"`       // 优惠活动落地页链接，若优惠活动说明已填，则必填
	ActivityStartTime int64  `json:"activity_start_time"` // 优惠开始时间（时间戳到秒），最多可早于提交时间 180 个自然日；若优惠活动说明已填，则必填
	ActivityType      string `json:"activity_type"`       // 优惠活动类型，默认值为优惠;不得出现特殊符号；可选文案：领券、红包、补贴、限免、特价、赠品、会员、拼团；如需新增，发邮件至 smartapp-yylzd@baidu.com ，邮件正文请提供 APP key 、小程序名称、页面标题、描述、落地页 path 、新增活动类型
}

type SubmitSkuRequestpriceInfo struct {
	OrgPrice      string `json:"org_price"`       // 优惠活动前的服务原价格
	OrgUnit       string `json:"org_unit"`        // 价格统一默认单位是「元」，提交价格字段时不必重复提交单位「元」，比如：98 元/月，价格字段提交 98、单位字段提交「月」；可选单位：小时、天、周、月、年、m² 、m³ 、次、件；支持新增；如需新增，发邮件至 ext_service_category@baidu.com ，邮件正文请提供 APP key 、小程序名称、页面标题、描述、落地页 path 、新增单位
	RangeMaxPrice string `json:"range_max_price"` // 当前服务的实际成交价格区间：最高价格
	RangeMinPrice string `json:"range_min_price"` // 当前服务的实际成交价格区间：最低价格
	RealPrice     string `json:"real_price"`      // 当前服务的实际成交价格，精确价格与价格区间不同时出现
}

type SubmitSkuRequestBody struct {
	ActivityInfo []SubmitSkuRequestactivityInfoItem `json:"activity_info"` // 优惠活动，详见：activity_info 字段说明
	ButtonName   string                             `json:"button_name"`   // 服务按钮内的文案，表达使用服务的动作；字数为两个/四个汉字，推荐文案如下：咨询/立即咨询、问诊/立即问诊、预约/立即预约、预订/立即预订、办理/立即办理、购买/立即购买、购票/立即购票、抢票/立即抢票、订票/立即订票、下单/立即下单、抢购/立即抢购、团购/立即团购、入住/立即入住、查询/立即查询、查看/立即查看、排队/立即排队、进店/立即进店、租车/立即租车、租房/立即租房、充值/立即充值、缴费/立即缴费；如需新增，发邮件至 ext_service_category@baidu.com ，邮件正文请提供 APP key 、小程序名称、页面标题、描述、落地页 path 、新增按钮文案
	Desc         string                             `json:"desc"`          // 服务简介，服务的文字解释说明，8~34 个字符(汉字占 2 字符)
	Images       []string                           `json:"images"`        // 封面图片链接，目前只须传 1 张图片 图片要求： 1.图片内容要求： 图片清晰、干净，不要出现令人不适的内容；不能出现严重影响用户理解的内容截断问题；图片无水印、二维码、logo； 2.相关性&一致性： 图片与标题、服务落地页内容相关、信息一致;单图片最大不能超 2M，只支持 JPG 或 PNG 格式 三种可选图片比例： a.正方形图片，比例 1:1，图片尺寸要求不低于 352 * 352 b.长方形横图，比例 16:9，图片尺寸不低于 1068 * 601 c.长方形竖图，比例 3:4，图片尺寸不低于 372 * 495
	Path         string                             `json:"path"`          // 智能小程序落地页链接
	PriceInfo    SubmitSkuRequestpriceInfo          `json:"price_info"`    // 服务价格，详见：price_info 字段说明
	Region       string                             `json:"region"`        // 服务地域，参考附录二，省市之间用英文中划线分割，多个地区之间用英文逗号分割
	Schema       string                             `json:"schema"`        // 行业扩展字段，trade_type=1001/1002/1004/5001/5003/7001/7002/7003/13001 时必须设置，详见：各个分类 schema 字段说明
	Tag          string                             `json:"tag"`           // 服务的属性、亮点；最多三个标签；每个标签字数不超过 5 个汉字；不得出现特殊符号；“官方标”不可提交；多个标签使用英文封号分割
	Title        string                             `json:"title"`         // 服务标题：描述服务是什么，12~30 个字符(汉字占 2 字符)
	TradeType    int64                              `json:"trade_type"`    // 服务类目编码，参考附录一
}

// 响应结构体

type SubmitSkuResponse struct {
	Data      string `json:"data"`       // 响应参数（如有提交失败的数据会返回失败的 path 和原因）
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// SubmitSku
func SubmitSku(params *SubmitSkuRequest) (string, error) {
	var (
		err        error
		defaultRet string
	)
	respData := &SubmitSkuResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeJSON).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/server/submit/sku")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	postData := make([]SubmitSkuRequestBody, 0, len(params.PostBody))
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
