package openapi

import (
	"fmt"
)

// ModifyMaterialRequest 请求结构体
type ModifyMaterialRequest struct {
	AccessToken   string // 接口调用凭证
	AppID         int64  // app_id
	ID            int64  // 物料 id ，添加物料时返回 id
	ImageURL      string // 小图片地址，长度不能超过 500 个字符，最小尺寸为 213*213，比例为 1：1，单图最大为 2M
	Title         string // 标题，需要描述完整，能够明确表示小程序或内容的主要信息点，不能全英文，6-30 个字。若选择相应垂类时，此字段只作为兜底展示字段
	Path          string // 智能小程序内页链接
	Category1Code string // 一级分类字段
	Category2Code string // 二级分类字段
	Desc          string // 4-17 个汉字。标题解释说明
	LabelAttr     string // 属性、特点。最多三个标签；每个标签字数不超过 5 个汉字,多个使用因为 / 隔开
	LabelDiscount string // 优惠信息，最多一个标签；每个标签字数不超过 7 个汉字
	ButtonName    string // 按钮文案，最多 4 个字
	BigImage      string // 封面图片链接（1 张，单图片最大 2M）大图模板要求最小尺寸 1068 x 601，比例为 16：9，单图最大为 2M
	VerticalImage string // 当选择小说/动漫，影视剧，电影票务，演出赛事时必填；（竖图 3：4）最低 213*284
	ExtJSON       string // 扩展信息
}

// 响应结构体

type ModifyMaterialResponse struct {
	Data      bool   `json:"data"`       // true：代表修改成功，false：代码修改失败
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// ModifyMaterial
func ModifyMaterial(params *ModifyMaterialRequest) (bool, error) {
	var (
		err        error
		defaultRet bool
	)
	respData := &ModifyMaterialResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/articlemount/material/modify")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("app_id", fmt.Sprintf("%v", params.AppID))
	client.AddPostParam("id", fmt.Sprintf("%v", params.ID))
	client.AddPostParam("imageUrl", params.ImageURL)
	client.AddPostParam("title", params.Title)
	client.AddPostParam("path", params.Path)
	client.AddPostParam("category1Code", params.Category1Code)
	client.AddPostParam("category2Code", params.Category2Code)
	client.AddPostParam("desc", params.Desc)
	client.AddPostParam("labelAttr", params.LabelAttr)
	client.AddPostParam("labelDiscount", params.LabelDiscount)
	client.AddPostParam("buttonName", params.ButtonName)
	client.AddPostParam("bigImage", params.BigImage)
	client.AddPostParam("verticalImage", params.VerticalImage)
	client.AddPostParam("extJson", params.ExtJSON)

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
