package pkg

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetPackageRequest 请求结构体
type GetPackageRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
}

// 响应结构体

type GetPackageResponsedataItem struct {
	CommitTime       string `json:"commit_time"`        // 提交时间
	Committer        string `json:"committer"`          // 提交人
	Msg              string `json:"msg"`                // 审核信息
	PackageID        int64  `json:"package_id"`         // 代码包 id
	Remark           string `json:"remark"`             // 备注
	RollbackVersion  string `json:"rollback_version"`   // 上一个线上版本的版本号
	Status           int64  `json:"status"`             // 状态
	UploadErrCode    int64  `json:"upload_err_code"`    // 上传状态为失败时，失败的错误码
	UploadErrMsg     string `json:"upload_err_msg"`     // 上传状态为失败时，失败的原因描述
	UploadStatus     int64  `json:"upload_status"`      // 上传状态
	UploadStatusDesc string `json:"upload_status_desc"` // 上传状态描述
	Version          string `json:"version"`            // 版本号
	VersionDesc      string `json:"version_desc"`       // 版本描述
}

type GetPackageResponse struct {
	Data      []GetPackageResponsedataItem `json:"data"`       // 响应参数
	Errno     int64                        `json:"errno"`      // 状态码
	ErrMsg    string                       `json:"msg"`        // 错误信息
	ErrorCode int64                        `json:"error_code"` // openapi 错误码
	ErrorMsg  string                       `json:"error_msg"`  // openapi 错误信息
}

// GetPackage
func GetPackage(params *GetPackageRequest) ([]GetPackageResponsedataItem, error) {
	var (
		err        error
		defaultRet []GetPackageResponsedataItem
	)
	respData := &GetPackageResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/package/get")
	client.AddGetParam("access_token", params.AccessToken)
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
	return respData.Data, nil
}
