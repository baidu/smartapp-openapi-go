package openapi

import (
	"fmt"
)

// SubmitResourceRequest 请求结构体
type SubmitResourceRequest struct {
	AccessToken string // 接口调用凭证
	AppID       int64  // app_id
	Body        string // 内容的介绍，最多 2000 字符
	Ext         string // 扩展信息（参考文档附录三）
	FeedSubType string // feed 二级分类（参考文档附录二）
	FeedType    string // feed 一级分类（参考文档附录二）
	Images      string // 封面图片链接，要求必须是 JSON 格式，最多 3 张，单图片最大不能超 2M，只支持 JPG 或 PNG 格式（jpeg 不支持），尺寸要求：宽不能低于 372px，且高不能低于 248px。重要提示：图片尺寸越大、清晰度越高、宽高比越接近3:2，越有助于降低不可用风险，促进分发。
	MappSubType string // 资源子类型（参考文档附录一）
	MappType    string // 资源类型（参考文档附录一）
	Path        string // 智能小程序落地页链接
	Tags        string // 资源标签，英文逗号分割，填写越准确详细可能带来更好的分发效果（最多 10 个，总长度最多 100 字）
	Title       string // 标题
}

// 响应结构体

type SubmitResourceResponse struct {
	Data      string `json:"data"`       // 响应参数
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// SubmitResource
func SubmitResource(params *SubmitResourceRequest) (string, error) {
	var (
		err        error
		defaultRet string
	)
	respData := &SubmitResourceResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/access/submitresource")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("app_id", fmt.Sprintf("%v", params.AppID))
	client.AddPostParam("body", params.Body)
	client.AddPostParam("ext", params.Ext)
	client.AddPostParam("feed_sub_type", params.FeedSubType)
	client.AddPostParam("feed_type", params.FeedType)
	client.AddPostParam("images", params.Images)
	client.AddPostParam("mapp_sub_type", params.MappSubType)
	client.AddPostParam("mapp_type", params.MappType)
	client.AddPostParam("path", params.Path)
	client.AddPostParam("tags", params.Tags)
	client.AddPostParam("title", params.Title)

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
