package msg

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetMsgRequest 请求结构体
type GetMsgRequest struct {
	AccessToken string      // 第三方平台的接口调用凭据
	Start       interface{} // 起始时间,默认值时间范围为近一天,起始时间不可超过一个月
	End         interface{} // 截止时间,默认值时间范围为为近一天,最大时间范围不可超过一周
	PushType    interface{} // 找回的推送类型 1:所有推送 2:失败推送
	Offset      interface{} // 分页参数(起始条数),默认值为0
	Count       interface{} // 分页参数(显示条数),默认值为100,最大值为100
	IDList      interface{} // 推送id,多个用逗号(,)拼接,若传入该字段,以上参数字段不影响结果
}

// 响应结构体

type GetMsgResponsedatalistItem struct {
	Content string `json:"content"` // 推送内容,加密后的结果
	ID      int64  `json:"id"`      // 推送id
	Status  int64  `json:"status"`  // 推送状态 0:待推送 1:推送成功 2:推送失败
}

type GetMsgResponsedata struct {
	Count    int64                        `json:"count"`    // 总条数
	List     []GetMsgResponsedatalistItem `json:"list"`     // 响应数据结构
	Page     int64                        `json:"page"`     // 当前页数
	PageSize int64                        `json:"pageSize"` // 单页条数
}

type GetMsgResponse struct {
	Data      GetMsgResponsedata `json:"data"`       // 响应参数
	Errno     int64              `json:"errno"`      // 状态码
	ErrMsg    string             `json:"msg"`        // 错误信息
	ErrorCode int64              `json:"error_code"` // openapi 错误码
	ErrorMsg  string             `json:"error_msg"`  // openapi 错误信息
}

// GetMsg
func GetMsg(params *GetMsgRequest) (*GetMsgResponsedata, error) {
	var (
		err        error
		defaultRet *GetMsgResponsedata
	)
	respData := &GetMsgResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pushmsg/getmsg")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("start", params.Start)
	client.AddGetParam("end", params.End)
	client.AddGetParam("push_type", params.PushType)
	client.AddGetParam("offset", params.Offset)
	client.AddGetParam("count", params.Count)
	client.AddGetParam("id_list", params.IDList)
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
