package openapi


// GetTemplateLibraryByIDRequest 请求结构体
type GetTemplateLibraryByIDRequest struct {
	AccessToken string // 接口调用凭证
	ID          string // 模板标题 id ，可通过接口获取，也可登录小程序后台查看获取
}

// 响应结构体

type GetTemplateLibraryByIDResponsedatakeywordListItem struct {
	Example   string `json:"example"`    // 关键词内容对应的示例
	KeywordID int64  `json:"keyword_id"` // 关键词 id ，添加模板时需要
	Name      string `json:"name"`       // 关键词内容
}

type GetTemplateLibraryByIDResponsedata struct {
	ID           string                                              `json:"id"`            // 模板标题 id
	KeywordCount int64                                               `json:"keyword_count"` // 关键字条数
	KeywordList  []GetTemplateLibraryByIDResponsedatakeywordListItem `json:"keyword_list"`  // 关键词数组，对象中包含关键词 id 、关键词内容 name 、关键词内容示例 example
	Title        string                                              `json:"title"`         // 模板标题
}

type GetTemplateLibraryByIDResponse struct {
	Data      GetTemplateLibraryByIDResponsedata `json:"data"`       // 响应对象
	Errno     int64                              `json:"errno"`      // 错误码
	ErrMsg    string                             `json:"msg"`        // 错误信息
	ErrorCode int64                              `json:"error_code"` // openapi 错误码
	ErrorMsg  string                             `json:"error_msg"`  // openapi 错误信息
}

// GetTemplateLibraryByID
func GetTemplateLibraryByID(params *GetTemplateLibraryByIDRequest) (*GetTemplateLibraryByIDResponsedata, error) {
	var (
		err        error
		defaultRet *GetTemplateLibraryByIDResponsedata
	)
	respData := &GetTemplateLibraryByIDResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/libraryget")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("id", params.ID)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)

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
