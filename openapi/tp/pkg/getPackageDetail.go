package pkg

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetPackageDetailRequest 请求结构体
type GetPackageDetailRequest struct {
	AccessToken string      // 授权小程序的接口调用凭据
	Type        interface{} // 小程序状态，不指定 package_id 的情况下默认是线上版本
	PackageID   interface{} // 代码包 id
}

// 响应结构体

type GetPackageDetailResponsedata struct {
	CommitTime  string `json:"commit_time"`  // 提交时间
	Committer   string `json:"committer"`    // 提交人
	Msg         string `json:"msg"`          // 审核信息
	PackageID   int64  `json:"package_id"`   // 代码包 id
	Remark      string `json:"remark"`       // 备注
	Status      int64  `json:"status"`       // 状态
	Version     string `json:"version"`      // 版本号
	VersionDesc string `json:"version_desc"` // 版本描述
}

type GetPackageDetailResponse struct {
	Data      GetPackageDetailResponsedata `json:"data"`       // 响应参数
	Errno     int64                        `json:"errno"`      // 状态码
	ErrMsg    string                       `json:"msg"`        // 错误信息
	ErrorCode int64                        `json:"error_code"` // openapi 错误码
	ErrorMsg  string                       `json:"error_msg"`  // openapi 错误信息
}

// GetPackageDetail
func GetPackageDetail(params *GetPackageDetailRequest) (*GetPackageDetailResponsedata, error) {
	var (
		err        error
		defaultRet *GetPackageDetailResponsedata
	)
	respData := &GetPackageDetailResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/package/getdetail")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("type", params.Type)
	client.AddGetParam("package_id", params.PackageID)
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
