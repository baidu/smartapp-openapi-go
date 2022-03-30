package msgtemplate

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetMsgTemplateLibraryKeyWordRequest 请求结构体
type GetMsgTemplateLibraryKeyWordRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	ID          string // 模板标题id，可通过接口获取，也可登录小程序后台查看获取
}

// 响应结构体

type GetMsgTemplateLibraryKeyWordResponsedatakeywordListItem struct {
	Example   string `json:"example"`    // 关键词内容对应的示例
	KeywordID int64  `json:"keyword_id"` // 关键词 id，添加模板时需要
	Name      string `json:"name"`       // 关键词内容
}

type GetMsgTemplateLibraryKeyWordResponsedata struct {
	ID           string                                                    `json:"id"`            // 模板标题 id
	KeywordCount int64                                                     `json:"keyword_count"` // 关键字条数
	KeywordList  []GetMsgTemplateLibraryKeyWordResponsedatakeywordListItem `json:"keyword_list"`  //
	Title        string                                                    `json:"title"`         // 模板标题
}

type GetMsgTemplateLibraryKeyWordResponse struct {
	Data      GetMsgTemplateLibraryKeyWordResponsedata `json:"data"`       // 响应参数
	Errno     int64                                    `json:"errno"`      // 状态码
	ErrMsg    string                                   `json:"msg"`        // 错误信息
	ErrorCode int64                                    `json:"error_code"` // openapi 错误码
	ErrorMsg  string                                   `json:"error_msg"`  // openapi 错误信息
}

// GetMsgTemplateLibraryKeyWord
func GetMsgTemplateLibraryKeyWord(params *GetMsgTemplateLibraryKeyWordRequest) (*GetMsgTemplateLibraryKeyWordResponsedata, error) {
	var (
		err        error
		defaultRet *GetMsgTemplateLibraryKeyWordResponsedata
	)
	respData := &GetMsgTemplateLibraryKeyWordResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/library/get")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("id", params.ID)
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
