package design

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
