package search

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAllSubChainRequest 请求结构体
type GetAllSubChainRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
}

// 响应结构体

type GetAllSubChainResponsedatasubChainInfoListItem struct {
	AppID        int64  `json:"app_id"`        // 小程序Id
	AppName      string `json:"app_name"`      // 小程序名称
	AuditDesc    string `json:"audit_desc"`    // 审核描述信息
	AuditTime    int64  `json:"audit_time"`    // 审核时间
	ChainDesc    string `json:"chain_desc"`    // 子链的描述信息
	ChainName    string `json:"chain_name"`    // 子链的名称
	ChainPath    string `json:"chain_path"`    // 子链的访问路径
	ChainRank    int64  `json:"chain_rank"`    // 展示顺序，当子链的状态为审核通过，值越小，越优先展示，展示数量由sub_chain_prior_number 控制，超过的子链为备选子链。新添加的子链该值为99，停用的子链该值为100。
	CreateTime   int64  `json:"create_time"`   // 创建时间
	CustomerID   int64  `json:"customer_id"`   // 主体id
	CustomerName string `json:"customer_name"` // 主体名称
	ID           int64  `json:"id"`            // 子链id
	PathMd5      string `json:"path_md5"`      //
	ShowStatus   int64  `json:"showStatus"`    // 展示类型。1: 优选 2: 备选。为null代表还未通过审核
	Status       int64  `json:"status"`        // 子链的状态。0：未知 1：审核中 2：审核通过 3：审核拒绝 4：撤回 5：删除 6：停用
	SubchainType int64  `json:"subchain_type"` // 子链类型。1: path类型子链 2:客服电话子链（要求小程序等级为 4 ）
	UpdateTime   int64  `json:"update_time"`   // 更新时间
}

type GetAllSubChainResponsedata struct {
	AppLevel            int64                                            `json:"app_level"`              // 小程序等级，等级为 3、4的小程序能展示六个单卡子链。
	SubChainInfoList    []GetAllSubChainResponsedatasubChainInfoListItem `json:"sub_chain_info_list"`    //
	SubChainPriorNumber int64                                            `json:"sub_chain_prior_number"` // 展示的单卡子链个数。
}

type GetAllSubChainResponse struct {
	Data      GetAllSubChainResponsedata `json:"data"`       // 响应参数
	Errno     int64                      `json:"errno"`      // 状态码
	ErrMsg    string                     `json:"msg"`        // 错误信息
	ErrorCode int64                      `json:"error_code"` // openapi 错误码
	ErrorMsg  string                     `json:"error_msg"`  // openapi 错误信息
}

// GetAllSubChain
func GetAllSubChain(params *GetAllSubChainRequest) (*GetAllSubChainResponsedata, error) {
	var (
		err        error
		defaultRet *GetAllSubChainResponsedata
	)
	respData := &GetAllSubChainResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/subchain/getall")
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
