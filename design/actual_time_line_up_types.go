package design

import (
	. "goa.design/goa/v3/dsl"
)

// 增量数据
var CanalDataResp = ResultType("CanalDataResp", func() {
	Description("数据库增量数据")
	TypeName("CanalDataResp")
	ContentType("application/json")

	Attributes(func() {
		Attribute("dataType", Int32, "数据类别", func() {
			Example(1035)
		})
		Attribute("dataInfo", Any, "数据", func() {
			Example(1035)
		})
		Required("dataType", "dataInfo")
	})

})
