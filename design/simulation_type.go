package design

import (
	. "goa.design/goa/v3/dsl"
)

var GetDataResp = ResultType("GetDataResp", func() {
	Description("获取模拟数据结果")
	TypeName("GetDataResp")
	ContentType("application/json")
	Attributes(func() {
		Attribute("key", String, "对应模块", func() {
			Example("对应模块")
		})
		Attribute("val", String, "对应值", func() {
			Example("对应值")
		})
		Attribute("isShowMock", Boolean, "是否显示配置项", func() {
			Example(true)
		})
		Attribute("orderBy", Int, "排序字段类型", func() {
			Description("1.通过评分正序 2.通过评分倒序")
			Example(1)
		})
		Attribute("orderTimeScope", Int, "排序时间范围", func() {
			Description("1.按年统计 2.按月统计")
			Example(1)
		})
		Required("key", "val", "isShowMock")
	})
})
