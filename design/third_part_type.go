package design

import (
	. "goa.design/goa/v3/dsl"
)

// 排号总览
var HallManagementResp = ResultType("HallManagementResp", func() {
	Description("大厅管理系统数据")
	TypeName("HallManagementResp")
	ContentType("application/json")
	Attributes(func() {
		Attribute("cardNum", String, "身份证", func() {
			Example("1035")
		})
		Attribute("name", String, "名字", func() {
			Example("小白")
		})
		Attribute("ouName", String, "部门名字", func() {
			Example("社保")
		})
		Required("cardNum", "name", "ouName")
	})
})

// 法人用户列表
var LegalPersonUserResp = ResultType("LegalPersonUserResp", func() {
	Description("法人用户列表")
	TypeName("LegalPersonUserResp")
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", Int, "用户Id", func() {
			Example(1035)
		})
		Attribute("name", String, "名字", func() {
			Example("小白")
		})
		Attribute("companies", ArrayOf(CompanyProfileResp), "公司", func() {

		})
		Required("id", "name", "companies")
	})
})

// 公司列表
var CompanyProfileResp = ResultType("CompanyProfileResp", func() {
	Description("公司简洁")
	TypeName("CompanyProfileResp")
	ContentType("application/json")

	Attributes(func() {
		Attribute("industry", Int, "行业Id", func() {
			Example(1036)
		})
		Attribute("name", String, "公司名字", func() {
			Example("青朵教育")
		})
		Attribute("userId", String, "法人名字", func() {
			Example("王大锤")
		})
		Required("industry", "name", "userId")
	})
})
