package openapi

// AddTemplateRequest 请求结构体
type AddTemplateRequest struct {
	AccessToken   string // 接口调用凭证
	ID            string // 模板标题 id
	KeywordIDList string // 模板关键词 id 列表，如[1,2,3]
}

// 响应结构体

type AddTemplateResponsedata struct {
	TemplateID string `json:"template_id"` // 添加至帐号下的模板 id ，发送小程序模板消息时所需
}

type AddTemplateResponse struct {
	Data      AddTemplateResponsedata `json:"data"`       // 响应对象
	Errno     int64                   `json:"errno"`      // 错误码
	ErrMsg    string                  `json:"msg"`        // 错误信息
	ErrorCode int64                   `json:"error_code"` // openapi 错误码
	ErrorMsg  string                  `json:"error_msg"`  // openapi 错误信息
}

// AddTemplate
func AddTemplate(params *AddTemplateRequest) (*AddTemplateResponsedata, error) {
	var (
		err        error
		defaultRet *AddTemplateResponsedata
	)
	respData := &AddTemplateResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/templateadd")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("id", params.ID)
	client.AddPostParam("keyword_id_list", params.KeywordIDList)

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
