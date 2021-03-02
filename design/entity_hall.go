package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("entity_hall", func() {
	Description("政务服务改革成效-实体大厅")
	Error("bad_request", ErrorResult)
	Error("internal_server_error", ErrorResult, func() {
		Fault()
	})

	HTTP(func() {
		Path("/entity_hall")
		Response("bad_request", StatusBadRequest)
		Response("internal_server_error", StatusInternalServerError)
	})

	Method("WaitLineOverview", func() {
		Description("排号总览")
		Meta("swagger:summary", "排号总览")
		Payload(func() {
			Attribute("regionCode", String, "行政区划代码", func() {
				Example("520100")
			})
			Attribute("startDate", String, "起始时间", func() {
				Example("2019-07-27")
			})
			Attribute("endDate", String, "结束时间", func() {
				Example("2020-07-27")
			})
			Required("regionCode")
		})
		Result(func() {
			Attribute("errcode", Int, "错误码", func() {
				Minimum(0)
				Maximum(999999)
				Example(0)
			})
			Attribute("errmsg", String, "错误消息", func() {
				Example("")
			})
			Attribute("data", WaitLineOverviewResp)
			Required("errcode", "errmsg", "data")
		})
		HTTP(func() {
			POST("/get_wait_line_overview")
			Response(StatusOK)
		})

	})
})
