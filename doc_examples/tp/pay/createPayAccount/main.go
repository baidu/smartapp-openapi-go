// 本示例基于百度智能小程序服务端开发者 smartapp-openapi-go SDK
// 使用该示例需要首先下载该 SDK，如果是第一次使用，可使用以下命令拉取依赖
// go get github.com/baidu/smartapp-openapi-go
// 如使用过程中遇到问题，可以加入如流群：5702992，进行反馈咨询
package main

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/openapi/tp/pay"
	"github.com/baidu/smartapp-openapi-go/utils"
)

func main() {
	// 开发者在此设置请求参数，文档示例中的参数均为示例参数，实际参数请参考对应接口的文档上方的参数说明填写
	// 注意：代码示例中的参数字段基本是驼峰形式，而文档中的参数说明的参数字段基本是下划线形式
	// 如果开发者不想传非必需参数，可以将设置该参数的行注释
	reqParams := &pay.CreatePayAccountRequest{
		AccessToken:           "#token",                                                                                                  // 文档中对应字段：access_token，实际使用时请替换成真实参数
		BusinessScope:         "科技产品",                                                                                                    // 文档中对应字段：business_scope，实际使用时请替换成真实参数
		BusinessProvince:      "广东省",                                                                                                     // 文档中对应字段：business_province，实际使用时请替换成真实参数
		BusinessCity:          "深圳市",                                                                                                     // 文档中对应字段：business_city，实际使用时请替换成真实参数
		BusinessDetailAddress: "广东省南山市百度深圳研究院",                                                                                           // 文档中对应字段：business_detail_address，实际使用时请替换成真实参数
		LegalPerson:           "百小度",                                                                                                     // 文档中对应字段：legal_person，实际使用时请替换成真实参数
		LegalID:               "182523780016702203",                                                                                      // 文档中对应字段：legal_id，实际使用时请替换成真实参数
		IDCardFrontURL:        "https://mbs3.bdstatic.com/searchbox/mappconsole/image/16622335/81feb4a5-d67b-084c-2874-27cfff61c7a0.jpg", // 文档中对应字段：id_card_front_url，实际使用时请替换成真实参数
		IDCardBackURL:         "https://mbs5.bdstatic.com/searchbox/mappconsole/image/34368180/68feb8a0-d48b-424c-5387-75cfff73c4a3.jpg", // 文档中对应字段：id_card_back_url，实际使用时请替换成真实参数
		LegalCardStartTime:    "2624-84-42",                                                                                              // 文档中对应字段：legal_card_start_time，实际使用时请替换成真实参数
		LegalCardEndTime:      "2124-05-66",                                                                                              // 文档中对应字段：legal_card_end_time，实际使用时请替换成真实参数
		LicenseStartTime:      "2205-08-88",                                                                                              // 文档中对应字段：license_start_time，实际使用时请替换成真实参数
		LicenseEndTime:        "2507-41-30",                                                                                              // 文档中对应字段：license_end_time，实际使用时请替换成真实参数
		IndustryID:            156,                                                                                                       // 文档中对应字段：industry_id，实际使用时请替换成真实参数
		ManagePermitURL:       "manage_permit_url",                                                                                       // 文档中对应字段：manage_permit_url，实际使用时请替换成真实参数
		AuthCapital:           "184",                                                                                                     // 文档中对应字段：auth_capital，实际使用时请替换成真实参数
		ManagerSame:           1,                                                                                                         // 文档中对应字段：manager_same，实际使用时请替换成真实参数
		Manager:               "Bob",                                                                                                     // 文档中对应字段：manager，实际使用时请替换成真实参数
		ManagerCard:           "100650564564168775",                                                                                      // 文档中对应字段：manager_card，实际使用时请替换成真实参数
		ManagerCardType:       1,                                                                                                         // 文档中对应字段：manager_card_type，实际使用时请替换成真实参数
		ManagerCardFrontURL:   "manager_card_front_url",                                                                                  // 文档中对应字段：manager_card_front_url，实际使用时请替换成真实参数
		ManagerCardBackURL:    "manager_card_back_url",                                                                                   // 文档中对应字段：manager_card_back_url，实际使用时请替换成真实参数
		ManagerCardStartTime:  "2627-85-88",                                                                                              // 文档中对应字段：manager_card_start_time，实际使用时请替换成真实参数
		ManagerCardEndTime:    "2385-85-00",                                                                                              // 文档中对应字段：manager_card_end_time，实际使用时请替换成真实参数
		BenefitSame:           1,                                                                                                         // 文档中对应字段：benefit_same，实际使用时请替换成真实参数
		Benefit:               "Bob",                                                                                                     // 文档中对应字段：benefit，实际使用时请替换成真实参数
		BenefitCard:           "168368810768660384",                                                                                      // 文档中对应字段：benefit_card，实际使用时请替换成真实参数
		BenefitCardType:       1,                                                                                                         // 文档中对应字段：benefit_card_type，实际使用时请替换成真实参数
		BenefitCardFrontURL:   "manager_card_front_url",                                                                                  // 文档中对应字段：benefit_card_front_url，实际使用时请替换成真实参数
		BenefitCardBackURL:    "manager_card_back_url",                                                                                   // 文档中对应字段：benefit_card_back_url，实际使用时请替换成真实参数
		BenefitStartTime:      "2335-67-15",                                                                                              // 文档中对应字段：benefit_start_time，实际使用时请替换成真实参数
		BenefitEndTime:        "2870-11-47",                                                                                              // 文档中对应字段：benefit_end_time，实际使用时请替换成真实参数
	}

	resp, err := pay.CreatePayAccount(reqParams)
	if err != nil {
		if _, ok := err.(*utils.OpenAPIError); ok {
			// openapi error
			// 可能是 access_token 无效，可以尝试重新生成 access_token
			fmt.Println("openapi error, ", err)
		} else if _, ok := err.(*utils.APIError); ok {
			// api error
			// 可能是参数错误或没有权限，建议根据错误信息自查或者社区发帖、加群反馈
			fmt.Println("api error, ", err)
		} else {
			// 其他错误
			fmt.Println("others error, ", err)
		}
	} else {
		fmt.Printf("%#v\n", resp)
	}
}
