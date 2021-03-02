package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("ActualTime", func() {
	Description("大厅排队办事实时图")

	Error("bad_request", ErrorResult)
	Error("internal_server_error", ErrorResult, func() {
		Fault()
	})
	HTTP(func() {
		Path("/actual_time")
		Response("bad_request", StatusBadRequest)
		Response("internal_server_error", StatusInternalServerError)
	})

	Method("GetActualTimeData", func() {
		Description("接收数据库监听得到的数据")
		Meta("swagger:summary", "接收数据库监听得到的数据")

		Payload(func() {
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
			Attribute("data", CanalDataResp)
			Required("errcode", "errmsg", "data")
		})

		HTTP(func() {
			GET("/get_actual_time_data")
			Response(StatusOK)
		})
	})

	Method("ReceiveThirdPartyPushData", func() {
		Description("接收第三方推送数据--大厅排队办事实时图基础数据")
		Meta("swagger:summary", "接收第三方推送数据")
		Payload(func() {
			Attribute("methodName", Int, "推送的具体方法", func() {
				Example(1)
			})
			Attribute("count", Int, "数据数量", func() {
				Example(21)
			})
			Attribute("data", Any, "第三方推送数据", func() {
				Example("")
			})

			Required("methodName", "count", "data")
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
			Attribute("result", String, "success", func() {
				Example("getDataSuccess")
			})
			Required("errcode", "errmsg", "result")
		})

		HTTP(func() {
			POST("/receive_third_party_push_data")
			Response(StatusOK)
		})
	})

})
