package template

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetTemplateListRequest 请求结构体
type GetTemplateListRequest struct {
	AccessToken string      // 第三方平台的接口调用凭据
	Page        interface{} // 页码（默认 1）
	PageSize    interface{} // 条数（默认 10）
}

// 响应结构体

type GetTemplateListResponsedatalistItem struct {
	CreateTime  int64  `json:"create_time"`  // 创建时间
	TemplateID  int64  `json:"template_id"`  // 模板 id
	UserDesc    string `json:"user_desc"`    // 模板描述名称
	UserVersion string `json:"user_version"` // 模板版本信息
	WebStatus   bool   `json:"web_status"`   // 是否支持 web 化，开发者工具编译版本 2.15.07 以上传的模板支持 web 化
}

type GetTemplateListResponsedata struct {
	Count int64                                 `json:"count"` // 数据数量
	List  []GetTemplateListResponsedatalistItem `json:"list"`  // 模版列表
}

type GetTemplateListResponse struct {
	Data      GetTemplateListResponsedata `json:"data"`       // 响应参数
	Errno     int64                       `json:"errno"`      // 状态码
	ErrMsg    string                      `json:"msg"`        // 错误信息
	ErrorCode int64                       `json:"error_code"` // openapi 错误码
	ErrorMsg  string                      `json:"error_msg"`  // openapi 错误信息
}

// GetTemplateList
func GetTemplateList(params *GetTemplateListRequest) (*GetTemplateListResponsedata, error) {
	var (
		err        error
		defaultRet *GetTemplateListResponsedata
	)
	respData := &GetTemplateListResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/gettemplatelist")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("page", params.Page)
	client.AddGetParam("page_size", params.PageSize)
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
