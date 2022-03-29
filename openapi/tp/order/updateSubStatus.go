package order

import (
	"encoding/json"
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// UpdateSubStatusRequest 请求结构体
type UpdateSubStatusRequest struct {
	AccessToken string                           // 授权小程序的接口调用凭据
	OpenID      string                           // 用户openId
	SceneID     string                           // 百度收银台分配的平台订单ID，通知支付状态接口返回的orderId
	SceneType   int64                            // 支付场景类型，开发者请默认传2
	Data        []UpdateSubStatusRequestDataItem //
}

type UpdateSubStatusRequestDataItemEXTSubsOrderItemsItem struct {
	SubOrderID string `json:"SubOrderID"` // 售后订单ID
	SubStatus  int64  `json:"SubStatus"`  // 自订单状态,枚举参照 【退换货枚举值】
}
type UpdateSubStatusRequestDataItemEXTSubsOrder struct {
	Items  []UpdateSubStatusRequestDataItemEXTSubsOrderItemsItem `json:"Items"`  //
	Status int64                                                 `json:"Status"` //
}
type UpdateSubStatusRequestDataItemEXT struct {
	SubsOrder UpdateSubStatusRequestDataItemEXTSubsOrder `json:"SubsOrder"` // 售后订单信息
}
type UpdateSubStatusRequestDataItem struct {
	BizAPPID   string                            `json:"BizAPPID"`   // 小程序的appKey
	CateID     int64                             `json:"CateID"`     // 2:订单种类-虚拟物品
	EXT        UpdateSubStatusRequestDataItemEXT `json:"EXT"`        // 拓展字段 根据资产的不同其结构也不固定 此处以订单为例
	ResourceID string                            `json:"ResourceID"` // 开发者接入的唯一订单ID
}

// 响应结构体

type UpdateSubStatusResponsedataItem struct {
	BizAppID     string `json:"biz_app_id"`    // POST请求参数中BizAPPID
	CateID       string `json:"cate_id"`       // POST请求参数中CateID
	ResourceID   string `json:"resource_id"`   // POST请求参数中ResourceID
	RowsAffected string `json:"rows_affected"` // 请求受影响行数(即请求是否成功 0为失败 非0为成功)
}

type UpdateSubStatusResponse struct {
	Data      []UpdateSubStatusResponsedataItem `json:"data"`       // 响应参数
	Errno     int64                             `json:"errno"`      // 状态码
	ErrMsg    string                            `json:"msg"`        // 错误信息
	ErrorCode int64                             `json:"error_code"` // openapi 错误码
	ErrorMsg  string                            `json:"error_msg"`  // openapi 错误信息
}

// UpdateSubStatus
func UpdateSubStatus(params *UpdateSubStatusRequest) ([]UpdateSubStatusResponsedataItem, error) {
	var (
		err        error
		defaultRet []UpdateSubStatusResponsedataItem
	)
	respData := &UpdateSubStatusResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeJSON).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/ordercenter/update/sub/status")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("open_id", params.OpenID)
	client.AddGetParam("scene_id", params.SceneID)
	client.AddGetParam("scene_type", fmt.Sprintf("%v", params.SceneType))
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	postData := map[string]interface{}{
		"Data": params.Data,
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
		return defaultRet, &utils.OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}

	if respData.Errno != 0 {
		return defaultRet, &utils.APIError{respData.Errno, respData.ErrMsg, respData}
	}
	return respData.Data, nil
}
