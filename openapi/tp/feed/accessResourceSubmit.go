package feed

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// AccessResourceSubmitRequest 请求结构体
type AccessResourceSubmitRequest struct {
	AccessToken string      // 授权小程序的接口调用凭据
	AppID       interface{} // 小程序 Id
	Body        string      // 消息体，物料的介绍
	Ext         interface{} // 扩展信息（JSON格式，参考附录三）
	FeedSubType interface{} // feed二级分类（参考附录二）
	FeedType    string      // feed一级分类（参考附录二）
	Images      interface{} // 封面图片链接（JSON格式）（最多3张，单图片最大2M） 建议尺寸：宽>=375 & 高>=250；建议比例 宽:高=1.5:1
	MappSubType int64       // 资源子类型（参考附录一）
	MappType    int64       // 资源类型（参考附录一）
	Path        string      // 智能小程序落地页链接
	Tags        interface{} // 资源标签，英文逗号分割，填写越准确详细可能带来更好的分发效果（最多10个，总长度最多100字）
	Title       string      // 标题
}

// 响应结构体

type AccessResourceSubmitResponse struct {
	Data      string `json:"data"`       // 响应参数
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// AccessResourceSubmit
func AccessResourceSubmit(params *AccessResourceSubmitRequest) (string, error) {
	var (
		err        error
		defaultRet string
	)
	respData := &AccessResourceSubmitResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/access/resource/submit")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("app_id", params.AppID)
	client.AddPostParam("body", params.Body)
	client.AddPostParam("ext", params.Ext)
	client.AddPostParam("feed_sub_type", params.FeedSubType)
	client.AddPostParam("feed_type", params.FeedType)
	client.AddPostParam("images", params.Images)
	client.AddPostParam("mapp_sub_type", fmt.Sprintf("%v", params.MappSubType))
	client.AddPostParam("mapp_type", fmt.Sprintf("%v", params.MappType))
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
		return defaultRet, &utils.OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}

	if respData.Errno != 0 {
		return defaultRet, &utils.APIError{respData.Errno, respData.ErrMsg, respData}
	}
	return respData.Data, nil
}
