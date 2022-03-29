package order

import (
	"encoding/json"
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// UpdateSubInfoRequest 请求结构体
type UpdateSubInfoRequest struct {
	AccessToken string                         // 授权小程序的接口调用凭据
	OpenID      string                         // 用户openId
	SceneID     string                         // 百度收银台分配的平台订单ID，通知支付状态接口返回的orderId
	SceneType   int64                          // 支付场景类型，开发者请默认传2
	Data        []UpdateSubInfoRequestDataItem //
}

type UpdateSubInfoRequestDataItemEXTSubsOrderItemsItemOrderDetail struct {
	AndroidSchema string `json:"AndroidSchema"` //
	H5Schema      string `json:"H5Schema"`      //
	IPhoneSchema  string `json:"IPhoneSchema"`  //
	Name          string `json:"Name"`          //
	Status        int64  `json:"Status"`        //
	SwanSchema    string `json:"SwanSchema"`    //
}
type UpdateSubInfoRequestDataItemEXTSubsOrderItemsItemRefundProductItem struct {
	Amount   int64  `json:"Amount"`   //
	ID       string `json:"ID"`       //
	Quantity int64  `json:"Quantity"` //
}
type UpdateSubInfoRequestDataItemEXTSubsOrderItemsItemRefund struct {
	Amount  int64                                                                `json:"Amount"`  //
	Product []UpdateSubInfoRequestDataItemEXTSubsOrderItemsItemRefundProductItem `json:"Product"` //
}
type UpdateSubInfoRequestDataItemEXTSubsOrderItemsItem struct {
	CTime       int64                                                        `json:"CTime"`       //
	MTime       int64                                                        `json:"MTime"`       //
	OrderDetail UpdateSubInfoRequestDataItemEXTSubsOrderItemsItemOrderDetail `json:"OrderDetail"` //
	OrderType   int64                                                        `json:"OrderType"`   //
	Refund      UpdateSubInfoRequestDataItemEXTSubsOrderItemsItemRefund      `json:"Refund"`      //
	SubOrderID  string                                                       `json:"SubOrderID"`  //
	SubStatus   int64                                                        `json:"SubStatus"`   //
}
type UpdateSubInfoRequestDataItemEXTSubsOrder struct {
	Items  []UpdateSubInfoRequestDataItemEXTSubsOrderItemsItem `json:"Items"`  //
	Status int64                                               `json:"Status"` //
}
type UpdateSubInfoRequestDataItemEXT struct {
	SubsOrder UpdateSubInfoRequestDataItemEXTSubsOrder `json:"SubsOrder"` //
}
type UpdateSubInfoRequestDataItem struct {
	BizAPPID   string                          `json:"BizAPPID"`   //
	CateID     int64                           `json:"CateID"`     //
	EXT        UpdateSubInfoRequestDataItemEXT `json:"EXT"`        //
	ResourceID string                          `json:"ResourceID"` //
}

// 响应结构体

type UpdateSubInfoResponsedataItem struct {
	BizAppID     string `json:"biz_app_id"`    // POST请求参数中BizAPPID
	CateID       string `json:"cate_id"`       // POST请求参数中CateID
	ResourceID   string `json:"resource_id"`   // POST请求参数中ResourceID
	RowsAffected string `json:"rows_affected"` // 请求受影响行数(即请求是否成功 0为失败 非0为成功)
}

type UpdateSubInfoResponse struct {
	Data      []UpdateSubInfoResponsedataItem `json:"data"`       // 响应参数
	Errno     int64                           `json:"errno"`      // 状态码
	ErrMsg    string                          `json:"msg"`        // 错误信息
	ErrorCode int64                           `json:"error_code"` // openapi 错误码
	ErrorMsg  string                          `json:"error_msg"`  // openapi 错误信息
}

// UpdateSubInfo
func UpdateSubInfo(params *UpdateSubInfoRequest) ([]UpdateSubInfoResponsedataItem, error) {
	var (
		err        error
		defaultRet []UpdateSubInfoResponsedataItem
	)
	respData := &UpdateSubInfoResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeJSON).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/ordercenter/update/sub/info")
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
