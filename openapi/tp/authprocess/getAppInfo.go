package authprocess

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAppInfoRequest 请求结构体
type GetAppInfoRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
}

// 响应结构体

type GetAppInfoResponsedataauditInfo struct {
	AuditAppDesc         string `json:"audit_app_desc"`          // 小程序介绍内容审核值
	AuditAppDescReason   string `json:"audit_app_desc_reason"`   // 小程序介绍内容失败原因
	AuditAppDescStatus   int64  `json:"audit_app_desc_status"`   // 小程序介绍内容审核状态
	AuditAppName         string `json:"audit_app_name"`          // 小程序名称审核值
	AuditAppNameReason   string `json:"audit_app_name_reason"`   // 小程序名称审核失败原因
	AuditAppNameStatus   int64  `json:"audit_app_name_status"`   // 小程序名称审核状态 1：审核中 3：审核失败
	AuditPhotoAddr       string `json:"audit_photo_addr"`        // 小程序头像审核值
	AuditPhotoAddrReason string `json:"audit_photo_addr_reason"` // 小程序头像审核失败原因
	AuditPhotoAddrStatus int64  `json:"audit_photo_addr_status"` // 小程序头像审核状态
}

type GetAppInfoResponsedataqualification struct {
	AdStatus int64  `json:"ad_status"` // 真实性认证状态 1：通过 3：失败
	AdType   int64  `json:"ad_type"`   // 真实性认证类型 -1：其他类型验证 0：未做真实性认证 1：对公验证 2：活体验证 23：法人人脸验证
	Name     string `json:"name"`      // 主体名称
	Satus    int64  `json:"satus"`     // 主体审核状态 1：通过 2：审核中 3：审核失败
	Type     int64  `json:"type"`      // 主体类型 1：个人 2：企业 3： 政府 4：媒体 5：其他 个人暂不开放
}

type GetAppInfoResponsedatamodifyCount struct {
	CategoryModifyQuota  int64 `json:"category_modify_quota"`  // 小程序类目总共的可修改次数
	CategoryModifyUsed   int64 `json:"category_modify_used"`   // 小程序类目已修改次数
	ImageModifyQuota     int64 `json:"image_modify_quota"`     // 小程序头像总共的可修改次数
	ImageModifyUsed      int64 `json:"image_modify_used"`      // 小程序头像已修改总次数
	NameModifyQuota      int64 `json:"name_modify_quota"`      // 小程序名称总共的可修改次数
	NameModifyUsed       int64 `json:"name_modify_used"`       // 小程序名称已修改次数
	SignatureModifyQuota int64 `json:"signature_modify_quota"` // 小程序简介总共的可修改次数
	SignatureModifyUsed  int64 `json:"signature_modify_used"`  // 小程序简介已修改次数
}

type GetAppInfoResponsedataappOfflineInfoItem struct {
	IllegalFields string `json:"illegal_fields"` // 强制下线原因 appName：名称 photoAddr：图片 appDesc：简介 当有多个时用逗号(,)连接, offlineReason为1或3时才有
	OfflineReason int64  `json:"offline_reason"` // 强制下线类型 1：基本信息强制下线 2：小程序代码包强制下线 3：基本信息和小程序代码包强制下线
}

type GetAppInfoResponsedataauthInfoItem struct {
	ScopeName string `json:"scope_name"` // 权限名称
	Type      int64  `json:"type"`       // 权限类型 ( 0：小程序纬度权限 1：账号纬度权限 )
}

type GetAppInfoResponsedatacategoryItemparent struct {
	CategoryDesc string `json:"category_desc"` // 父类目描述
	CategoryID   int64  `json:"category_id"`   // 父类目 id
	CategoryName string `json:"category_name"` // 父类目名称
}

type GetAppInfoResponsedatacategoryItem struct {
	AuditStatus  int64                                    `json:"audit_status"`  // 类目审核状态 1：审核中 2：审核成功 3：审核失败
	CategoryDesc string                                   `json:"category_desc"` // 类目描述
	CategoryID   int64                                    `json:"category_id"`   // 类目 id
	CategoryName string                                   `json:"category_name"` // 类目名称
	Parent       GetAppInfoResponsedatacategoryItemparent `json:"parent"`        // 父类目
	Reason       string                                   `json:"reason"`        // 审核失败原因
}

type GetAppInfoResponsedataannualReviewInfo struct {
	AnnualReviewOverdueTime int64 `json:"annual_review_overdue_time"` // 年审过期时间
	AnnualReviewStatus      int64 `json:"annual_review_status"`       // 年审状态 1：正常 2：待年审 3： 年审过期
}

type GetAppInfoResponsedata struct {
	AnnualReviewInfo GetAppInfoResponsedataannualReviewInfo     `json:"annual_review_info"` // 小程序年审相关信息
	AppDesc          string                                     `json:"app_desc"`           // 小程序的介绍内容
	AppID            int64                                      `json:"app_id"`             // 小程序的 appid
	AppKey           string                                     `json:"app_key"`            // 小程序的 appkey
	AppName          string                                     `json:"app_name"`           // 小程序的名称
	AppOfflineInfo   []GetAppInfoResponsedataappOfflineInfoItem `json:"app_offline_info"`   // 小程序强制下线相关信息
	AuditInfo        GetAppInfoResponsedataauditInfo            `json:"audit_info"`         // 基本信息审核状态 包括名称，图标，介绍内容的审核状态 只有审核中和审核失败会展示
	AuthInfo         []GetAppInfoResponsedataauthInfoItem       `json:"auth_info"`          // 小程序权限集合信息
	Category         []GetAppInfoResponsedatacategoryItem       `json:"category"`           // 小程序类目信息
	MinSwanVersion   string                                     `json:"min_swan_version"`   // 开发者工具最低版本
	ModifyCount      GetAppInfoResponsedatamodifyCount          `json:"modify_count"`       // 小程序基本信息修改次数信息
	PhotoAddr        string                                     `json:"photo_addr"`         // 小程序图标
	Qualification    GetAppInfoResponsedataqualification        `json:"qualification"`      // 小程序账号对应的主体信息
	Status           int64                                      `json:"status"`             // 小程序的状态 -1：代表封禁 1：代表正常 2：代表审核中 4：代表暂停服务 5：强制下线 6：限时整改 7：流量下线
	WebStatus        int64                                      `json:"web_status"`         // 小程序的web化开关状态 0：未开启 1：开启 2：关闭
}

type GetAppInfoResponse struct {
	Data      GetAppInfoResponsedata `json:"data"`       // 响应参数
	Errno     int64                  `json:"errno"`      // 状态码
	ErrMsg    string                 `json:"msg"`        // 错误信息
	ErrorCode int64                  `json:"error_code"` // openapi 错误码
	ErrorMsg  string                 `json:"error_msg"`  // openapi 错误信息
}

// GetAppInfo
func GetAppInfo(params *GetAppInfoRequest) (*GetAppInfoResponsedata, error) {
	var (
		err        error
		defaultRet *GetAppInfoResponsedata
	)
	respData := &GetAppInfoResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/app/info")
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
	return &respData.Data, nil
}
