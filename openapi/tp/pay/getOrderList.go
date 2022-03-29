package pay

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetOrderListRequest 请求结构体
type GetOrderListRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	Status      string // 订单状态。 all：全部 、 1：待付款 、 2：已付款 、 3：已消费 、 4：退款中 、 5：已退款 、 6：退款失败 、7：已取消
	StartTime   int64  // 起始时间戳，10位时间戳
	EndTime     int64  // 起始时间戳，10位时间戳
	CurrentPage int64  // 当前页数。起始为 1
	PageSize    int64  // 分页。每页数量
}

// 响应结构体

type GetOrderListResponsedatadataItem struct {
	Channel              string `json:"channel"`                // 支付渠道。如：网银支付、微信支付
	CreateTime           string `json:"create_time"`            // 订单创建时间。格式为 yyyy-MM-dd HH:mm:ss，如：2019-06-07 00:00:00
	DealID               string `json:"deal_id"`                // 百度收银台的财务结算凭证。详见电商技术平台术语
	DealTitle            string `json:"deal_title"`             // 商品名。订单的名称
	ID                   string `json:"id"`                     // 百度平台订单ID。
	OrderID              string `json:"order_id"`               // 平台订单号。百度平台订单ID
	PayMoney             string `json:"pay_money"`              // 支付金额(分)。扣除各种优惠后用户还需要支付的金额
	Phone                string `json:"phone"`                  // 用户手机号。
	SegmentRefundedMoney string `json:"segment_refunded_money"` // 退款金额(分)。
	SubStatus            string `json:"sub_status"`             // 订单状态 1:待付款 2:已付款 3:已消费 4:退款中 5:已退款 6:退款失败 7:已取消
	TotalMoney           string `json:"total_money"`            // 订单金额(分)。订单的实际金额
	TpOrderID            string `json:"tp_order_id"`            // 第三方订单号。业务方唯一订单号
	UpdateTime           string `json:"update_time"`            // 订单完成时间。格式为 yyyy-MM-dd HH:mm:ss，如：2019-06-07 00:00:00
}

type GetOrderListResponsedata struct {
	CurrentPage int64 `json:"current_page"` // 当前页码

	Data       []GetOrderListResponsedatadataItem `json:"data"`        // 数据详情
	PageSize   int64                              `json:"page_size"`   // 每页数量
	TotalCount int64                              `json:"total_count"` // 数据总数
	TotalMoney string                             `json:"total_money"` // 订单金额(分)。订单的实际金额
	TotalPage  int64                              `json:"total_page"`  // 总页码数
}

type GetOrderListResponse struct {
	Data      GetOrderListResponsedata `json:"data"`       //
	Errno     int64                    `json:"errno"`      //
	ErrMsg    string                   `json:"msg"`        // 错误信息
	ErrorCode int64                    `json:"error_code"` // openapi 错误码
	ErrorMsg  string                   `json:"error_msg"`  // openapi 错误信息
}

// GetOrderList
func GetOrderList(params *GetOrderListRequest) (*GetOrderListResponsedata, error) {
	var (
		err        error
		defaultRet *GetOrderListResponsedata
	)
	respData := &GetOrderListResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/orderlist")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("status", params.Status)
	client.AddGetParam("start_time", fmt.Sprintf("%v", params.StartTime))
	client.AddGetParam("end_time", fmt.Sprintf("%v", params.EndTime))
	client.AddGetParam("current_page", fmt.Sprintf("%v", params.CurrentPage))
	client.AddGetParam("page_size", fmt.Sprintf("%v", params.PageSize))
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
