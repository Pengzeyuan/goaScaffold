package design

import (
	. "goa.design/goa/v3/dsl"
)

var FourDo = ResultType("FourDo", func() {
	Description("四个办统计")
	TypeName("FourDo")
	ContentType("application/json")

	Attributes(func() {
		Attribute("year", Int, "年份")
		Attribute("count", Int, "数量")

		Required("year", "count")
	})
})

var FourDoCountResp = ResultType("FourDoCountResp", func() {
	Description("四个办统计信息")
	TypeName("FourDoCountResp")
	ContentType("application/json")
	Attributes(func() {
		Attribute("immediateInfoCount", ArrayOf(FourDo, func() {
			View("default")
		}))
		Attribute("onlineInfoCount", ArrayOf(FourDo, func() {
			View("default")
		}))
		Attribute("nearbyInfoCount", ArrayOf(FourDo, func() {
			View("default")
		}))
		Attribute("onceInfoCount", ArrayOf(FourDo, func() {
			View("default")
		}))
		Required("immediateInfoCount", "onlineInfoCount", "nearbyInfoCount", "onceInfoCount")
	})
})

var ReformOfAdministrativeResp = ResultType("ReformOfAdministrativeResp", func() {
	Description("行政审批制度改革事项详情")
	TypeName("ReformOfAdministrativeResp")
	ContentType("application/json")
	Attributes(func() {

		Attribute("splitCount", Int32, "今年事项大项拆分", func() {
			Example(48)
		})
		Attribute("pastSplitCount", Int32, "去年事项大项拆分", func() {
			Example(65)
		})
		Attribute("splitRate", Float32, "事项拆分比率", func() {
			Example(32.5)
		})

		Required("splitCount", "pastSplitCount", "splitRate")
	})

	View("reform_administrative", func() {
		Attribute("splitCount")
		Attribute("pastSplitCount")
		Attribute("splitRate")
	})
})

var CrowdRunsLittleResp = ResultType("CrowdRunsLittleResp", func() {
	Description("群众少跑腿返回值")
	TypeName("CrowdRunsLittleResp")
	ContentType("application/json")
	Attributes(func() {
		Attribute("mattersAccounted", ArrayOf(MatterNumber, func() {
			View("default")
		}), "最多跑一次的事项不同年份数据")

		Attribute("mattersAccounted", ArrayOf(MatterNumber, func() {
			View("default")
		}), "总事项不同年份数据")

		Attribute("mattersAccountedProportion", Float32, "最多跑一次事项占比提升比例", func() {
			Example(0.8)
		})
		Required("mattersAccounted", "mattersAccountedProportion")
	})
})

var MatterNumber = ResultType("MatterNumber", func() {
	Description("事项数")
	TypeName("MatterNumber")
	ContentType("application/json")
	Attributes(func() {
		Attribute("beforeAscension", Int32, "提升前", func() {
			Example(120)
		})
		Attribute("AfterAscension", Int32, "提升后", func() {
			Example(120)
		})
		Required("beforeAscension", "AfterAscension")
	})
})
