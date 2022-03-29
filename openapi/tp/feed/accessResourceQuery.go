package feed

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// AccessResourceQueryRequest 请求结构体
type AccessResourceQueryRequest struct {
	AccessToken string      // 授权小程序的接口调用凭据
	Begin       interface{} // 开始时间 (默认值前一天0点)
	End         interface{} // 结束时间 (默认值今天0点)
	ImageType   interface{} // 图片类型(1:封面图片 2:问答用户头像 3:动态图片 4:feed物料图片)
	PageNo      interface{} // 页数(分页参数,第几页,默认值(1)
	PageSize    interface{} // 单页展示数据量(分页参数,默认值(10)
	Status      interface{} // 状态（0: 全部 1: 审核中 2: 审核失败 3: 投放中 4: 已删除），默认值为0
	Title       interface{} // 标题
}

// 响应结构体

type AccessResourceQueryResponsedatalistItem struct {
	AppID      int64  `json:"app_id"`      // 小程序id
	AuditTime  string `json:"audit_time"`  // 审核时间
	CreateTime string `json:"create_time"` // 提交时间
	Rid        string `json:"rid"`         // 物料id
	Source     string `json:"source"`      // 物料来源
	Status     string `json:"status"`      // 状态
	Title      string `json:"title"`       // 标题
}

type AccessResourceQueryResponsedata struct {
	Count    int64                                     `json:"count"`    // 数据总量
	List     []AccessResourceQueryResponsedatalistItem `json:"list"`     // 数据内容
	Page     int64                                     `json:"page"`     // 页数
	PageSize int64                                     `json:"pageSize"` // 单页数量
}

type AccessResourceQueryResponse struct {
	Data      AccessResourceQueryResponsedata `json:"data"`       // 响应参数
	Errno     int64                           `json:"errno"`      // 状态码
	ErrMsg    string                          `json:"msg"`        // 错误信息
	ErrorCode int64                           `json:"error_code"` // openapi 错误码
	ErrorMsg  string                          `json:"error_msg"`  // openapi 错误信息
}

// AccessResourceQuery
func AccessResourceQuery(params *AccessResourceQueryRequest) (*AccessResourceQueryResponsedata, error) {
	var (
		err        error
		defaultRet *AccessResourceQueryResponsedata
	)
	respData := &AccessResourceQueryResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/access/resource/query")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("begin", params.Begin)
	client.AddGetParam("end", params.End)
	client.AddGetParam("image_type", params.ImageType)
	client.AddGetParam("page_no", params.PageNo)
	client.AddGetParam("page_size", params.PageSize)
	client.AddGetParam("status", params.Status)
	client.AddGetParam("title", params.Title)
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
