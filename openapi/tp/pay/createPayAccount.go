package pay

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// CreatePayAccountRequest 请求结构体
type CreatePayAccountRequest struct {
	AccessToken           string      // 授权小程序的接口调用凭据
	BusinessScope         string      // 经营范围。营业执照上经营范围，长度 2-2000 个字符
	BusinessProvince      string      // 经营范围所在省。
	BusinessCity          string      // 经营范围所在市。
	BusinessDetailAddress string      // 经营地址详情。
	LegalPerson           string      // 法人姓名，即身份证姓名（汉字，限制在 1 - 25 位）。
	LegalID               string      // 身份证号（长度限定为18位）。
	IDCardFrontURL        string      // 身份证正面地址(必须是我们服务上传生成的图片URL，参见图片上传)
	IDCardBackURL         string      // 身份证反面地址(必须是我们服务上传生成的图片URL，参见图片上传)
	LegalCardStartTime    string      // 法人身份证开始时间 例：2020-12-31
	LegalCardEndTime      string      // 法人身份证结束时间 例：2020-12-31，长期有效传 9999-12-31
	LicenseStartTime      string      // 营业执照证开始时间 例：2020-12-31，长期有效传营业执照核准日期
	LicenseEndTime        string      // 营业执照结束时间 例：2020-12-31，长期有效传 9999-12-31
	IndustryID            int64       // 行业id 见查询行业id列表接口
	ManagePermitURL       string      // 若行业id需要资质，资质图片地址
	AuthCapital           string      // 注册资本
	ManagerSame           int64       // 经营控股人是否与法人一致，0 - 不一致；1 - 一致。如果不一致则相关信息必传否则不传
	Manager               interface{} // 最大股东姓名
	ManagerCard           interface{} // 最大股东身份证
	ManagerCardType       interface{} // 最大股东身份证类型
	ManagerCardFrontURL   interface{} // 最大股东身份证正面地址
	ManagerCardBackURL    interface{} // 最大股东身份证反面地址
	ManagerCardStartTime  interface{} // 最大股东证件开始时间
	ManagerCardEndTime    interface{} // 最大股东证件结束时间
	BenefitSame           int64       // 受益人是否与法人一致，0 - 不一致；1 - 一致。如果不一致则相关信息必传否则不传
	Benefit               interface{} // 受益人姓名
	BenefitCard           interface{} // 受益人身份证
	BenefitCardType       interface{} // 受益人身份证类型
	BenefitCardFrontURL   interface{} // 受益人身份证正面地址
	BenefitCardBackURL    interface{} // 受益人身份证反面地址
	BenefitStartTime      interface{} // 受益人证件开始时间
	BenefitEndTime        interface{} // 受益人证件结束时间
}

// 响应结构体

type CreatePayAccountResponsedata struct {
}

type CreatePayAccountResponse struct {
	Data      CreatePayAccountResponsedata `json:"data"`       // 响应参数
	Errno     int64                        `json:"errno"`      // 状态码
	ErrMsg    string                       `json:"msg"`        // 错误信息
	ErrorCode int64                        `json:"error_code"` // openapi 错误码
	ErrorMsg  string                       `json:"error_msg"`  // openapi 错误信息
}

// CreatePayAccount
func CreatePayAccount(params *CreatePayAccountRequest) (*CreatePayAccountResponsedata, error) {
	var (
		err        error
		defaultRet *CreatePayAccountResponsedata
	)
	respData := &CreatePayAccountResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/account/create")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("business_scope", params.BusinessScope)
	client.AddPostParam("business_province", params.BusinessProvince)
	client.AddPostParam("business_city", params.BusinessCity)
	client.AddPostParam("business_detail_address", params.BusinessDetailAddress)
	client.AddPostParam("legal_person", params.LegalPerson)
	client.AddPostParam("legal_id", params.LegalID)
	client.AddPostParam("id_card_front_url", params.IDCardFrontURL)
	client.AddPostParam("id_card_back_url", params.IDCardBackURL)
	client.AddPostParam("legal_card_start_time", params.LegalCardStartTime)
	client.AddPostParam("legal_card_end_time", params.LegalCardEndTime)
	client.AddPostParam("license_start_time", params.LicenseStartTime)
	client.AddPostParam("license_end_time", params.LicenseEndTime)
	client.AddPostParam("industry_id", fmt.Sprintf("%v", params.IndustryID))
	client.AddPostParam("manage_permit_url", params.ManagePermitURL)
	client.AddPostParam("auth_capital", params.AuthCapital)
	client.AddPostParam("manager_same", fmt.Sprintf("%v", params.ManagerSame))
	client.AddPostParam("manager", params.Manager)
	client.AddPostParam("manager_card", params.ManagerCard)
	client.AddPostParam("manager_card_type", params.ManagerCardType)
	client.AddPostParam("manager_card_front_url", params.ManagerCardFrontURL)
	client.AddPostParam("manager_card_back_url", params.ManagerCardBackURL)
	client.AddPostParam("manager_card_start_time", params.ManagerCardStartTime)
	client.AddPostParam("manager_card_end_time", params.ManagerCardEndTime)
	client.AddPostParam("benefit_same", fmt.Sprintf("%v", params.BenefitSame))
	client.AddPostParam("benefit", params.Benefit)
	client.AddPostParam("benefit_card", params.BenefitCard)
	client.AddPostParam("benefit_card_type", params.BenefitCardType)
	client.AddPostParam("benefit_card_front_url", params.BenefitCardFrontURL)
	client.AddPostParam("benefit_card_back_url", params.BenefitCardBackURL)
	client.AddPostParam("benefit_start_time", params.BenefitStartTime)
	client.AddPostParam("benefit_end_time", params.BenefitEndTime)

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
