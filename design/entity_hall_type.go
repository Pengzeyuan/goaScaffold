package design

import (
	. "goa.design/goa/v3/dsl"
)

// 排号总览
var WaitLineOverviewResp = ResultType("WaitLineOverviewResp", func() {
	Description("排号总览")
	TypeName("WaitLineOverviewResp")
	ContentType("application/json")
	Attributes(func() {
		Attribute("todayDQ", Int32, "累计排号数", func() {
			Example(1035)
		})
		Attribute("cumulativeDQ", Int32, "累计办件量", func() {
			Example(89362)
		})
		Required("todayDQ", "cumulativeDQ")
	})
})
